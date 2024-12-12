package services

import (
	"backend-challenge/models"
	"backend-challenge/repositories"
	"fmt"

	"github.com/google/uuid"
)

type CreateFarmService struct {
	repository repositories.CreateFarmRepositoryInterface
}

func NewCreateFarmService(repository repositories.CreateFarmRepositoryInterface) *CreateFarmService {
	return &CreateFarmService{repository: repository}
}

func (s *CreateFarmService) CreateFarm(farm *models.Farm) error {
	farm.ID = uuid.New()

	for i := range farm.Productions {
		farm.Productions[i].ID = uuid.New()
		farm.Productions[i].FarmID = farm.ID
	}

	err := s.repository.CreateFarm(farm)
	if err != nil {
		return fmt.Errorf("failed to create farm: %v", err)
	}

	return nil
}
