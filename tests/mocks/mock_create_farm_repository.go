package mocks

import (
	"backend-challenge/models"

	"github.com/stretchr/testify/mock"
)

type MockCreateFarmRepository struct {
	mock.Mock
}

func (m *MockCreateFarmRepository) CreateFarm(farm *models.Farm) error {
	args := m.Called(farm)
	return args.Error(0)
}
