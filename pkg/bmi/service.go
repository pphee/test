package bmi

import (
	"context"
	"fmt"
	"github.com/anush008/fastembed-go"
	"math/rand"
	"time"
)

type BmiService struct {
	repo Repository
}

func NewServices(bq Repository) *BmiService {
	return &BmiService{
		repo: bq,
	}
}

func init() {
	rand.Seed(time.Now().UnixNano())
}

func generateRandomPerson() BMI {
	names := []string{"Alice", "Bob", "Charlie", "David", "Eve", "Frank", "Grace", "Hank", "Ivy", "Jack"}
	name := names[rand.Intn(len(names))]
	age := rand.Intn(60) + 18
	height := float64(rand.Intn(50) + 150)
	weight := float64(rand.Intn(50) + 50)
	bmiValue := weight / ((height / 100) * (height / 100))
	return BMI{
		Name:   name,
		Age:    int64(age),
		Height: height,
		Weight: weight,
		Value:  bmiValue,
	}
}

func (u *BmiService) StoreBMI(ctx context.Context) ([]*BMI, error) {
	model, err := fastembed.NewFlagEmbedding(nil)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return nil, fmt.Errorf("failed to initialize embedding model: %w", err)
	}
	defer func() {
		if destroyErr := model.Destroy(); destroyErr != nil {
			fmt.Printf("Error destroying model: %v\n", destroyErr)
		}
	}()

	var bmis []*BMI
	var batch []*BMI
	const totalRecords = 1000
	const totalEmbeddings = 100
	const batchSize = 5

	for i := 0; i < totalRecords; i++ {
		person := generateRandomPerson()
		bmi := &BMI{
			Name:         person.Name,
			Age:          person.Age,
			Height:       person.Height,
			Weight:       person.Weight,
			Value:        person.Value,
			CreatedAt:    time.Now(),
			RecordNumber: i + 1,
		}
		bmi.Category, bmi.Risk = CalculateBMICategoryAndRisks(person.Value)

		if i < totalEmbeddings {
			text := fmt.Sprintf("Name: %s, Age: %d, BMI: %.2f, Height: %.2f, Weight: %.2f",
				person.Name, person.Age, person.Value, person.Height, person.Weight)

			embeddings, err := model.Embed([]string{text}, batchSize)
			if err != nil {
				fmt.Printf("Error: Failed to generate embedding for record %d: %v\n", bmi.RecordNumber, err)
				bmi.Vector = []float32{0.0, 0.0, 0.0}
			} else if len(embeddings) == 0 || len(embeddings[0]) < 3 {
				fmt.Printf("Warning: Empty or invalid embeddings for record %d, embeddings: %v\n", bmi.RecordNumber, embeddings)
				bmi.Vector = []float32{0.0, 0.0, 0.0}
			} else {
				bmi.Vector = []float32{embeddings[0][0], embeddings[0][1], embeddings[0][2]}
			}
			bmi.VectorGenerated = true
		} else {
			bmi.Vector = []float32{0.0, 0.0, 0.0}
			bmi.VectorGenerated = true
		}

		bmis = append(bmis, bmi)
		batch = append(batch, bmi)

		if len(batch) == batchSize || i == totalRecords-1 {
			err := u.repo.StoreBatch(ctx, batch)
			if err != nil {
				fmt.Printf("Error: Failed to store batch ending with record %d: %v\n", bmi.RecordNumber, err)
			} else {
				fmt.Printf("Successfully stored batch ending with record %d\n", bmi.RecordNumber)
			}
			batch = []*BMI{}
		}

		time.Sleep(10 * time.Millisecond)
	}

	return bmis, nil
}

func CalculateBMICategoryAndRisks(bmiValue float64) (category, risk string) {
	switch {
	case bmiValue < 18.5:
		category = "Underweight"
		risk = "Low"
	case bmiValue < 25:
		category = "Normal weight"
		risk = "Average"
	case bmiValue < 30:
		category = "Overweight"
		risk = "Increased"
	default:
		category = "Obesity"
		risk = "High"
	}
	return category, risk
}

func (s *BmiService) QueryBMI(ctx context.Context, queryVector []float32) ([]*BMIIV, error) {
	scoredPoints, err := s.repo.Query(ctx, queryVector)
	if err != nil {
		return nil, err
	}

	var results []*BMIIV

	for _, sp := range scoredPoints {
		if sp == nil || sp.Payload == nil {
			continue
		}

		height := 0.0
		if h, ok := sp.Payload["height"]; ok {
			height = h.GetDoubleValue()
		}

		name := ""
		if n, ok := sp.Payload["name"]; ok {
			name = n.GetStringValue()
		}

		age := 0
		if a, ok := sp.Payload["age"]; ok {
			if a.GetIntegerValue() != 0 {
				age = int(a.GetIntegerValue())
			} else if a.GetDoubleValue() != 0 {
				age = int(a.GetDoubleValue())
			}
		}

		weight := 0.0
		if w, ok := sp.Payload["weight"]; ok {
			weight = w.GetDoubleValue()
		}

		value := 0.0
		if v, ok := sp.Payload["value"]; ok {
			value = v.GetDoubleValue()
		}

		category := ""
		if c, ok := sp.Payload["category"]; ok {
			category = c.GetStringValue()
		}

		risk := ""
		if r, ok := sp.Payload["risk"]; ok {
			risk = r.GetStringValue()
		}

		var createdAt time.Time
		if v, ok := sp.Payload["created_at"]; ok {
			if createdAtString := v.GetStringValue(); createdAtString != "" {
				if parsedTime, err := time.Parse(time.RFC3339, createdAtString); err == nil {
					createdAt = parsedTime
				}
			}
		}

		var uuidId string

		if sp.Id != nil {
			if tmpUuidId := sp.Id.GetUuid(); tmpUuidId != "" {
				uuidId = tmpUuidId
			}
		}

		bmi := &BMIIV{
			Uuid:      uuidId,
			Name:      name,
			Age:       int64(age),
			Height:    height,
			Weight:    weight,
			Value:     value,
			Category:  category,
			Risk:      risk,
			CreatedAt: createdAt,
		}

		results = append(results, bmi)
	}

	return results, nil
}
