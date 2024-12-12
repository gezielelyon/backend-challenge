package mocks

import (
	"backend-challenge/models"

	"github.com/stretchr/testify/mock"
)

type MockCreateFarmService struct {
	mock.Mock
}

func (m *MockCreateFarmService) CreateFarm(farm *models.Farm) error {
	args := m.Called(farm)
	return args.Error(0)
}
