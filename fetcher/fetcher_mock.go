package fetcher

import (
	"github.com/stretchr/testify/mock"
	"github.com/syned13/clinics-api/models"
)

// FetcherMock entity for mocking Fetcher methods
type FetcherMock struct {
	mock.Mock
}

// FetchClinics mockws the FetchClinics fetcher function
func (m *FetcherMock) FetchClinics(clinicType models.ClinicType) ([]models.Clinic, error) {
	args := m.Called(clinicType)

	return args.Get(0).([]models.Clinic), args.Error(1)
}
