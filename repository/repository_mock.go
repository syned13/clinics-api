package repository

import (
	"github.com/stretchr/testify/mock"
	"github.com/syned13/clinics-api/models"
)

// RepositoryMock entity for mocking the Repository methods
type RepositoryMock struct {
	mock.Mock
}

// GetClinics mocks the GetClinics repository methods
func (m *RepositoryMock) GetClinics(name string, state string, from, to string) ([]models.Clinic, error) {
	args := m.Called(name, state, from, to)

	return args.Get(0).([]models.Clinic), args.Error(1)
}

func (m *RepositoryMock) UpdateClinics(clinics []models.Clinic) error {
	args := m.Called(clinics)

	return args.Error(1)
}
