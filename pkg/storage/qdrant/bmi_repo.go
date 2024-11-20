package qdrant

import (
	client "github.com/qdrant/go-client/qdrant"
)

type BMIRepository struct {
	client         *client.Client
	collectionName string
}
