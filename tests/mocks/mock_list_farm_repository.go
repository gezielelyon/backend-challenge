package mocks

import (
	"backend-challenge/models"

	"github.com/stretchr/testify/mock"
)

type MockListFarmRepository struct {
	mock.Mock
}

func (m *MockListFarmRepository) ListFarms(page, limit int, cropType, landArea string) ([]models.Farm, error) {
	args := m.Called(page, limit, cropType, landArea)
	return args.Get(0).([]models.Farm), args.Error(1)
}
