package bmi

import (
	"context"
	client "github.com/qdrant/go-client/qdrant"
)

type Repository interface {
	StoreBatch(ctx context.Context, bmis []*BMI) error
	Query(ctx context.Context, queryVector []float32) ([]*client.ScoredPoint, error)
}
