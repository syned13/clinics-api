package service

import (
	"github.com/stretchr/testify/mock"
	"github.com/syned13/clinics-api/models"
)

// MockService entity for mocking the service methods
type MockService struct {
	mock.Mock
}

// GetClinics mocks the GetClinics service method
func (m *MockService) GetClinics(name string, state string, from, to string) ([]models.Clinic, error) {
	args := m.Called(name, state, from, to)

	return args.Get(0).([]models.Clinic), args.Error(1)
}

// UpdateClinics mocks the UpdateClinics service method
func (m *MockService) UpdateClinics(clinics []models.Clinic) error {
	args := m.Called(clinics)

	return args.Error(0)
}

// UpdateClinicsFromAPI mocks the UpdateClinicsFromAPI service method
func (m *MockService) UpdateClinicsFromAPI() error {
	args := m.Called()

	return args.Error(0)
}
