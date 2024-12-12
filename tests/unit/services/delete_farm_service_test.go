package services

import (
	"backend-challenge/services"
	"backend-challenge/tests/mocks"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDeleteFarmService_DeleteFarm(t *testing.T) {
	mockRepo := new(mocks.MockDeleteFarmRepository)

	service := services.NewDeleteFarmService(mockRepo)

	farmID := "12345"

	mockRepo.On("DeleteFarm", farmID).Return(nil)

	err := service.DeleteFarm(farmID)

	assert.NoError(t, err)
	mockRepo.AssertCalled(t, "DeleteFarm", farmID)
}
