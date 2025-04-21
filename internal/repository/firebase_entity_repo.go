package repository

import (
	"cloud.google.com/go/firestore"
	"context"
	"github.com/google/uuid"
	"go-dash-api/internal/model"
	"go-dash-api/internal/util"
	"log"
)

func NewFirebaseEntityRepository(db *firestore.Client) *FirebaseEntityRepository {
	return &FirebaseEntityRepository{db: db}
}

type FirebaseEntityRepository struct {
	db *firestore.Client
}

func (repo *FirebaseEntityRepository) GetByID(collection, id string) (*model.Entity, error) {
	fireEntity := new(model.Entity)
	obj, err := repo.db.Collection(collection).Doc(id).Get(context.Background())
	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	fireEntity.Collection = collection
	fireEntity.ID = obj.Ref.ID
	if err := obj.DataTo(&fireEntity.Props); err != nil {
		log.Fatal(err)
		return nil, err
	}
	return fireEntity, nil
}

func (repo *FirebaseEntityRepository) GetAll(collection string, page, limit int) ([]*model.Entity, error) {
	users := make([]*model.Entity, 0)
	objList, err := repo.db.Collection(collection).Limit(limit).Offset((page - 1) * limit).Documents(context.Background()).GetAll()
	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	for _, obj := range objList {
		firebaseEntity := new(model.Entity)
		firebaseEntity.Collection = collection
		firebaseEntity.ID = obj.Ref.ID
		if err := obj.DataTo(&firebaseEntity.Props); err != nil {
			log.Fatal(err)
			return nil, err
		}
		users = append(users, firebaseEntity)
	}
	return users, nil
}

func (repo *FirebaseEntityRepository) Create(entity *model.Entity) error {
	entity.ID = uuid.New().String()
	if _, err := repo.db.Collection(entity.Collection).Doc(entity.ID).Set(context.Background(), entity.Props); err != nil {
		log.Fatal(err)
		return err
	}
	return nil
}

func (repo *FirebaseEntityRepository) Update(entity *model.Entity) error {
	mapEntity := util.ToFirestoreUpdateArray(entity.Props)
	if _, err := repo.db.Collection(entity.Collection).Doc(entity.ID).Update(context.Background(), mapEntity); err != nil {
		log.Fatal(err)
		return err
	}
	return nil
}

func (repo *FirebaseEntityRepository) Delete(collection, id string) error {
	if _, err := repo.db.Collection(collection).Doc(id).Delete(context.Background()); err != nil {
		log.Fatal(err)
		return err
	}
	return nil
}
