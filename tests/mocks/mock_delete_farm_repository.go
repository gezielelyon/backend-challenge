package mocks

import (
	"github.com/stretchr/testify/mock"
)

type MockDeleteFarmRepository struct {
	mock.Mock
}

func (m *MockDeleteFarmRepository) DeleteFarm(farmID string) error {
	args := m.Called(farmID)
	return args.Error(0)
}
