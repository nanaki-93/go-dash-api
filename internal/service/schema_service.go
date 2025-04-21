package service

import (
	"go-dash-api/internal/model"
	"go-dash-api/internal/repository"
)

type SchemaService struct {
	repo repository.SchemaRepository
}

func NewSchemaService(repo repository.SchemaRepository) *SchemaService {
	return &SchemaService{repo: repo}
}

func (s *SchemaService) GetSchema(name string) (*model.Schema, error) {
	return s.repo.GetByName(name)
}

func (s *SchemaService) GetSchemas(page, limit int) ([]*model.Schema, error) {
	return s.repo.GetAll(page, limit)
}

func (s *SchemaService) AddSchema(entity *model.Schema) error {
	return s.repo.Create(entity)
}
func (s *SchemaService) UpdateSchema(entity *model.Schema) error {
	return s.repo.Update(entity)
}
func (s *SchemaService) DeleteSchema(id string) error {
	return s.repo.Delete(id)
}
