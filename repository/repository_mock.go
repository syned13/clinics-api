package repository

import (
	"github.com/stretchr/testify/mock"
	"github.com/syned13/clinics-api/models"
)

// MockRepository entity for mocking the Repository methods
type MockRepository struct {
	mock.Mock
}

// GetClinics mocks the GetClinics repository method
func (m *MockRepository) GetClinics(name string, state string, from, to string) ([]models.Clinic, error) {
	args := m.Called(name, state, from, to)

	return args.Get(0).([]models.Clinic), args.Error(1)
}

// UpdateClinics mocks the UpdateClinics repository method
func (m *MockRepository) UpdateClinics(clinics []models.Clinic) error {
	args := m.Called(clinics)

	return args.Error(1)
}
