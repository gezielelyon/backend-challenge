package services

import (
	"backend-challenge/repositories"
	"fmt"
)

type DeleteFarmService struct {
	repository repositories.DeleteFarmRepositoryInterface
}

func NewDeleteFarmService(repository repositories.DeleteFarmRepositoryInterface) *DeleteFarmService {
	return &DeleteFarmService{repository: repository}
}

func (s *DeleteFarmService) DeleteFarm(farmID string) error {
	err := s.repository.DeleteFarm(farmID)
	if err != nil {
		return fmt.Errorf("failed to delete farm: %v", err)
	}

	return nil
}
