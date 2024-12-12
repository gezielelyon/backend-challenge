package controllers

import (
	"backend-challenge/controllers"
	"backend-challenge/models"
	"backend-challenge/tests/mocks"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestListFarmController_Handle(t *testing.T) {
	mockService := new(mocks.MockListFarmService)

	controller := controllers.NewListFarmController(mockService)

	farms := []models.Farm{
		{
			ID:            uuid.Must(uuid.Parse("22c31a2b-6aa6-45ad-9bf7-c3c9c08acaba")),
			Name:          "Farm 1",
			LandArea:      100,
			UnitOfMeasure: "acre",
			Address:       "123 Test Street",
		},
		{
			ID:            uuid.Must(uuid.Parse("22c31a2b-6aa6-45ad-9bf7-c3c9c08acaba")),
			Name:          "Farm 2",
			LandArea:      200,
			UnitOfMeasure: "hectare",
			Address:       "456 Test Avenue",
		},
	}

	mockService.On("ListFarms", 1, 10, "", "").Return(farms, nil)

	req := httptest.NewRequest(http.MethodGet, "/farms?page=1&limit=10", nil)
	w := httptest.NewRecorder()

	controller.Handle(w, req)

	assert.Equal(t, http.StatusOK, w.Code)

	var response []models.Farm
	err := json.NewDecoder(w.Body).Decode(&response)
	assert.NoError(t, err)
	assert.Equal(t, 2, len(response))
	assert.Equal(t, "Farm 1", response[0].Name)
	mockService.AssertCalled(t, "ListFarms", 1, 10, "", "")
}

func TestListFarmController_Handle_InvalidPage(t *testing.T) {
	mockService := new(mocks.MockListFarmService)

	controller := controllers.NewListFarmController(mockService)

	req := httptest.NewRequest(http.MethodGet, "/farms?page=invalid&limit=10", nil)
	w := httptest.NewRecorder()

	controller.Handle(w, req)

	assert.Equal(t, http.StatusBadRequest, w.Code)
}
