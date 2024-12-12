package repositories

import "backend-challenge/models"

type CreateFarmRepositoryInterface interface {
	CreateFarm(farm *models.Farm) error
}

type ListFarmRepositoryInterface interface {
	ListFarms(page, limit int, cropType, landArea string) ([]models.Farm, error)
}

type DeleteFarmRepositoryInterface interface {
	DeleteFarm(farmID string) error
}
