package services

import (
	"backend-challenge/models"
	"backend-challenge/repositories"
)

type ListFarmService struct {
	repository repositories.ListFarmRepositoryInterface
}

func NewListFarmService(repository repositories.ListFarmRepositoryInterface) *ListFarmService {
	return &ListFarmService{repository: repository}
}

func (s *ListFarmService) ListFarms(page, limit int, cropType, landArea string) ([]models.Farm, error) {
	return s.repository.ListFarms(page, limit, cropType, landArea)
}
