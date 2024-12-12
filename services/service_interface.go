package services

import "backend-challenge/models"

type CreateFarmServiceInterface interface {
	CreateFarm(farm *models.Farm) error
}

type ListFarmServiceInterface interface {
	ListFarms(page, limit int, cropType, landArea string) ([]models.Farm, error)
}

type DeleteFarmServiceInterface interface {
	DeleteFarm(farmID string) error
}
