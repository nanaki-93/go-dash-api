package repository

import (
	"go-dash-api/internal/model"
)

type EntityRepository interface {
	GetByID(collection, id string) (*model.Entity, error)
	GetAll(collection string, page, limit int) ([]*model.Entity, error)
	Create(user *model.Entity) error
	Update(user *model.Entity) error
	Delete(collection string, id string) error
}

type SchemaRepository interface {
	GetByName(id string) (*model.Schema, error)
	GetAll(page, limit int) ([]*model.Schema, error)
	Create(user *model.Schema) error
	Update(user *model.Schema) error
	Delete(id string) error
}
