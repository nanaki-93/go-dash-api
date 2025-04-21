package repository

import (
	"cloud.google.com/go/firestore"
	"context"
	"github.com/google/uuid"
	"go-dash-api/internal/model"
	"log"
)

func NewFirebaseSchemaRepository(db *firestore.Client) *FirebaseSchemaRepository {
	return &FirebaseSchemaRepository{db: db}
}

type FirebaseSchemaRepository struct {
	db *firestore.Client
}

func (repo *FirebaseSchemaRepository) GetByName(name string) (*model.Schema, error) {
	fireSchema := new(model.Schema)
	obj, err := repo.db.Collection("schema").Where("Name", "==", name).Documents(context.Background()).Next()
	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	if err := obj.DataTo(&fireSchema); err != nil {
		log.Fatal(err)
		return nil, err
	}
	return fireSchema, nil
}

func (repo *FirebaseSchemaRepository) GetAll(page, limit int) ([]*model.Schema, error) {
	schemas := make([]*model.Schema, 0)
	objList, err := repo.db.Collection("schema").Limit(limit).Offset((page - 1) * limit).Documents(context.Background()).GetAll()
	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	for _, obj := range objList {
		fireSchema := new(model.Schema)
		if err := obj.DataTo(&fireSchema); err != nil {
			log.Fatal(err)
			return nil, err
		}
		schemas = append(schemas, fireSchema)
	}
	return schemas, nil
}

func (repo *FirebaseSchemaRepository) Create(schema *model.Schema) error {
	schema.Id = uuid.New().String()
	if _, err := repo.db.Collection("schema").Doc(schema.Id).Set(context.Background(), schema); err != nil {
		log.Fatal("repo - Create:", err)
		return err
	}
	return nil
}

func (repo *FirebaseSchemaRepository) Update(schema *model.Schema) error {
	mapSchema := make([]firestore.Update, 1)
	mapSchema[0] = firestore.Update{Path: "Structure", Value: schema.Structure}

	if _, err := repo.db.Collection("schema").Doc(schema.Id).Update(context.Background(), mapSchema); err != nil {
		log.Fatal(err)
		return err
	}
	return nil
}

func (repo *FirebaseSchemaRepository) Delete(id string) error {
	if _, err := repo.db.Collection("schema").Doc(id).Delete(context.Background()); err != nil {
		log.Fatal(err)
		return err
	}
	return nil
}
