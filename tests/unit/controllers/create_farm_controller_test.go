package controllers

import (
	"backend-challenge/controllers"
	"backend-challenge/models"
	"backend-challenge/tests/mocks"
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateFarmController_Handle(t *testing.T) {
	mockService := new(mocks.MockCreateFarmService)

	controller := controllers.NewCreateFarmController(mockService)

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

	mockService.On("CreateFarm", farm).Return(nil)

	payload, _ := json.Marshal(farm)

	req := httptest.NewRequest(http.MethodPost, "/farms", bytes.NewBuffer(payload))
	w := httptest.NewRecorder()

	controller.Handle(w, req)

	assert.Equal(t, http.StatusCreated, w.Code)

	var response models.Farm
	err := json.NewDecoder(w.Body).Decode(&response)
	assert.NoError(t, err)
	assert.NotEqual(t, response.ID, "")
	assert.Equal(t, response.Name, "Test Farm")
	mockService.AssertCalled(t, "CreateFarm", farm)
}

func TestCreateFarmController_Handle_InvalidRequest(t *testing.T) {
	mockService := new(mocks.MockCreateFarmService)

	controller := controllers.NewCreateFarmController(mockService)

	invalidFarm := &models.Farm{
		Name:          "",
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

	payload, _ := json.Marshal(invalidFarm)

	req := httptest.NewRequest(http.MethodPost, "/farms", bytes.NewBuffer(payload))
	w := httptest.NewRecorder()

	controller.Handle(w, req)

	assert.Equal(t, http.StatusBadRequest, w.Code)
}
