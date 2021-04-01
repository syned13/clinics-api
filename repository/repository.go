package repository

import (
	"context"

	"github.com/syned13/clinics-api/models"
)

type Repository interface {
	GetClinics(ctx context.Context, name string, state string, from, to string) ([]models.Clinic, error)
	UpdateClinics(clinics []models.Clinic) error
}

type clinicsRepository struct {
	sortedClinics  []models.Clinic            // sorted by opening time
	clinicsByState map[string][]models.Clinic // each list is sorted by
	index          index
}

// UpdateClinics updates the store with the new clinics
func (c clinicsRepository) UpdateClinics(clinics []models.Clinic) error {
	// sortedClinics := make([]models.Clinic, len(clinics))
	// clinicsByState := map[string]models.Clinic{}

	// sort.

	return nil
}
