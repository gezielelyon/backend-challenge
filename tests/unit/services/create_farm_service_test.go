package services

import (
	"backend-challenge/models"
	"backend-challenge/services"
	"backend-challenge/tests/mocks"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestCreateFarmService_CreateFarm(t *testing.T) {
	mockRepo := new(mocks.MockCreateFarmRepository)

	service := services.NewCreateFarmService(mockRepo)

	farm := &models.Farm{
		Name:          "Test Farm",
		LandArea:      100,
		UnitOfMeasure: "acre",
		Address:       "123 Test Street",
		Productions: []models.Production{
			{
				CropType:    "CORN",
				IsIrrigated: true,
				IsInsured:   false,
			},
		},
	}

	mockRepo.On("CreateFarm", farm).Return(nil)

	err := service.CreateFarm(farm)

	assert.NoError(t, err)
	assert.NotEqual(t, uuid.Nil, farm.ID)
	assert.NotEqual(t, uuid.Nil, farm.Productions[0].ID)
	mockRepo.AssertCalled(t, "CreateFarm", farm)
}
