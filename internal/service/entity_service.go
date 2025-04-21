package service

import (
	"go-dash-api/internal/model"
	"go-dash-api/internal/repository"
)

type EntityService struct {
	repo repository.EntityRepository
}

func NewEntityService(repo repository.EntityRepository) *EntityService {
	return &EntityService{repo: repo}
}

func (s *EntityService) GetEntity(collection, id string) (*model.Entity, error) {
	return s.repo.GetByID(collection, id)
}

func (s *EntityService) GetEntities(collection string, page, limit int) ([]*model.Entity, error) {
	return s.repo.GetAll(collection, page, limit)
}

func (s *EntityService) AddEntity(entity *model.Entity) error {
	return s.repo.Create(entity)
}
func (s *EntityService) UpdateEntity(entity *model.Entity) error {
	return s.repo.Update(entity)
}
func (s *EntityService) DeleteEntity(collection, id string) error {
	return s.repo.Delete(collection, id)
}
