package qdrant

import (
	"context"
	"fmt"
	"github.com/pphee/test/pkg/bmi"
	client "github.com/qdrant/go-client/qdrant"
	uuid "github.com/satori/go.uuid"
	"math"
	"time"
)

type BMIRepository struct {
	client         *client.Client
	collectionName string
}

func NewBMIRepository(host, apiKey, collectionName string, qdrantPort int) (*BMIRepository, error) {
	if len(host) >= 8 && (host[:7] == "http://" || host[:8] == "https://") {
		return nil, fmt.Errorf("host should not include scheme (http:// or https://), got: %s", host)
	}

	qdrantClient, err := client.NewClient(&client.Config{
		Host:   host,
		APIKey: apiKey,
		Port:   qdrantPort,
		UseTLS: qdrantPort == 443,
	})
	if err != nil {
		return nil, fmt.Errorf("failed to create Qdrant client: %w", err)
	}

	return &BMIRepository{
		client:         qdrantClient,
		collectionName: collectionName,
	}, nil
}

func (r *BMIRepository) CreateCollection(ctx context.Context) error {
	err := r.client.CreateCollection(ctx, &client.CreateCollection{
		CollectionName: r.collectionName,
		VectorsConfig: client.NewVectorsConfig(&client.VectorParams{
			Size:     3,
			Distance: client.Distance_Cosine,
		}),
	})
	if err != nil {
		return fmt.Errorf("failed to create collection: %w", err)
	}
	return nil
}

func (r *BMIRepository) StoreBatch(ctx context.Context, bmis []*bmi.BMI) error {
	points := make([]*client.PointStruct, 0, len(bmis))

	for _, bmi := range bmis {
		pointID := uuid.NewV4().String()
		payload := client.NewValueMap(map[string]any{
			"name":       bmi.Name,
			"age":        bmi.Age,
			"height":     bmi.Height,
			"weight":     bmi.Weight,
			"value":      bmi.Value,
			"category":   bmi.Category,
			"risk":       bmi.Risk,
			"created_at": bmi.CreatedAt.Format(time.RFC3339),
		})

		if !bmi.VectorGenerated {
			fmt.Printf("warning: bmi.Vector is not set for record %d, assigning zero vector\n", bmi.RecordNumber)
			bmi.Vector = []float32{0.0, 0.0, 0.0}
		} else if len(bmi.Vector) != 3 {
			fmt.Printf("warning: invalid bmi.Vector for record %d, assigning zero vector\n", bmi.RecordNumber)
			bmi.Vector = []float32{0.0, 0.0, 0.0}
		}

		normalizedVector := make([]float32, len(bmi.Vector))
		for j := range normalizedVector {
			normalizedVector[j] = bmi.Vector[j] / 100.0
		}

		vector := client.NewVectors(normalizedVector...)
		if vector == nil {
			fmt.Printf("error: failed to create vector for record %d\n", bmi.RecordNumber)
			continue
		}

		point := &client.PointStruct{
			Id:      client.NewID(pointID),
			Vectors: vector,
			Payload: payload,
		}

		points = append(points, point)
	}

	if len(points) == 0 {
		return fmt.Errorf("no valid points to upsert")
	}

	ctx, cancel := context.WithTimeout(ctx, 300*time.Second)
	defer cancel()

	for retries := 0; retries < 3; retries++ {
		retryCtx, retryCancel := context.WithTimeout(ctx, 300*time.Second)
		defer retryCancel()

		_, err := r.client.Upsert(retryCtx, &client.UpsertPoints{
			CollectionName: r.collectionName,
			Points:         points,
		})
		if err != nil {
			fmt.Printf("error: failed to upsert batch, retrying (%d/3): %v\n", retries+1, err)
			time.Sleep(time.Duration(math.Pow(2, float64(retries))) * time.Second)
			continue
		}
		return nil
	}

	return fmt.Errorf("failed to upsert batch after 3 retries")
}

func (r *BMIRepository) Query(ctx context.Context, queryVector []float32) ([]*client.ScoredPoint, error) {
	limit := uint64(10)

	response, err := r.client.Query(ctx, &client.QueryPoints{
		CollectionName: r.collectionName,
		Query:          client.NewQuery(queryVector...),
		Limit:          &limit,
		WithPayload:    client.NewWithPayload(true),
	})
	if err != nil {
		fmt.Printf("error: %v\n", err)
		return nil, fmt.Errorf("failed to query points: %w", err)
	}

	return response, nil
}
