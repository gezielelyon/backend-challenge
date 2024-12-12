package services

import (
	"backend-challenge/models"
	"backend-challenge/services"
	"backend-challenge/tests/mocks"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestListFarmService_ListFarms(t *testing.T) {
	mockRepo := &mocks.MockListFarmRepository{}
	mockService := services.NewListFarmService(mockRepo)

	expectedFarms := []models.Farm{
		{
			ID:            uuid.Must(uuid.Parse("22c31a2b-6aa6-45ad-9bf7-c3c9c08acaba")),
			Name:          "Farm 1",
			LandArea:      100,
			UnitOfMeasure: "acre",
			Address:       "123 Test St",
			Productions:   []models.Production{},
		},
	}

	mockRepo.On("ListFarms", 1, 10, "CORN", "100").Return(expectedFarms, nil)

	farms, err := mockService.ListFarms(1, 10, "CORN", "100")

	assert.NoError(t, err)
	assert.Equal(t, expectedFarms, farms)
	mockRepo.AssertExpectations(t)
}
