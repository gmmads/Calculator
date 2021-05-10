package repository

import (
	"context"
	"log"

	"cloud.google.com/go/firestore"
	"github.com/gmmads/Calculator/entity"
	"google.golang.org/api/iterator"
)

type repo struct{}

func NewFirestoreRepository() CalculationRepository {
	return &repo{}
}

const (
	projectId      string = "calculator-37a6b"
	collectionName string = "posts"
)

func (*repo) Save(calculation *entity.Calculation) (*entity.Calculation, error) {
	ctx := context.Background()
	client, err := firestore.NewClient(ctx, projectId)
	if err != nil {
		log.Fatalf("Failed to create a firestore client: %v", err)
		return nil, err
	}
	defer client.Close()
	_, _, err = client.Collection(collectionName).Add(ctx, map[string]interface{}{
		"expr":   calculation.Expr,
		"result": calculation.Result,
	})

	if err != nil {
		log.Fatalf("Failed adding to history: %v", err)
		return nil, err
	}
	return calculation, nil
}

func (*repo) FindAll() ([]entity.Calculation, error) {
	ctx := context.Background()
	client, err := firestore.NewClient(ctx, projectId)
	if err != nil {
		log.Fatalf("Failed to create a firestore client: %v", err)
		return nil, err
	}
	defer client.Close()

	var calculations []entity.Calculation
	iter := client.Collection(collectionName).Documents(ctx)
	for {
		doc, err := iter.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			log.Fatalf("Failed to iterate the list of calculations: %v", err)
			return nil, err
		}
		calculation := entity.Calculation{
			Expr:   doc.Data()["expr"].(string),
			Result: doc.Data()["result"].(float64),
		}
		calculations = append(calculations, calculation)
	}
	return calculations, nil
}
