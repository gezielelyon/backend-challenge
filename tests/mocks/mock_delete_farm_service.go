package mocks

import (
	"github.com/stretchr/testify/mock"
)

type MockDeleteFarmService struct {
	mock.Mock
}

func (m *MockDeleteFarmService) DeleteFarm(farmID string) error {
	args := m.Called(farmID)
	return args.Error(0)
}
