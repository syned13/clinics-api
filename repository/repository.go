package repository

import (
	"sort"

	"github.com/syned13/clinics-api/models"
	"github.com/syned13/clinics-api/utils"
)

type Repository interface {
	GetClinics(name string, state string, from, to string) ([]models.Clinic, error)
	UpdateClinics(clinics []models.Clinic) error
}

type clinicsRepository struct {
	sortedClinics []models.Clinic // sorted by opening time
	index         index
}

// NewClinicRepository returns a clinic repository entity
func NewClinicRepository() Repository {
	return &clinicsRepository{}
}

func (c clinicsRepository) GetClinics(name string, state string, from, to string) ([]models.Clinic, error) {
	clinics := c.sortedClinics

	if name != "" {
		byNameIndexes := c.index.Search(name)
		if len(byNameIndexes) != 0 {
			clinics = c.filterByIndexes(clinics, byNameIndexes)
		}
	}

	if state != "" {
		clinics = filterByState(clinics, state)
	}

	if from != "" && to != "" {
		clinics = filterByAvailability(clinics, from, to)
	}

	return clinics, nil
}

func isValidTime(from, to string) bool {
	return utils.ValidateHour(from) && utils.ValidateHour(to)
}

func filterByAvailability(clinics []models.Clinic, from, to string) []models.Clinic {
	output := make([]models.Clinic, 0, len(clinics))

	for _, clinic := range clinics {
		if !isValidTime(clinic.Availability.From, clinic.Availability.To) {
			continue
		}

		if clinic.Availability.From <= from && clinic.Availability.To >= to {
			output = append(output, clinic)
		}
	}

	return output
}

func filterByState(clinics []models.Clinic, state string) []models.Clinic {
	output := make([]models.Clinic, 0, len(clinics))

	for _, clinic := range clinics {
		if clinic.Statename == state {
			output = append(output, clinic)
		}
	}

	return output
}

func (c clinicsRepository) filterByIndexes(clinics []models.Clinic, indexes []int) []models.Clinic {
	output := make([]models.Clinic, 0, len(clinics))

	for _, value := range indexes {
		if value >= 0 && value < len(clinics) {
			output = append(output, clinics[value])
		}
	}

	return output
}

// UpdateClinics updates the store with the new clinics
func (c *clinicsRepository) UpdateClinics(clinics []models.Clinic) error {
	sortedClinics := make([]models.Clinic, len(clinics))
	copy(sortedClinics, clinics)

	sort.Slice(sortedClinics, func(i, j int) bool {
		return sortedClinics[i].Availability.From < sortedClinics[j].Availability.From
	})

	clinicNames := takeClinicNames(sortedClinics)

	c.index = NewIndex(clinicNames)
	c.sortedClinics = sortedClinics

	return nil
}

func takeClinicNames(clinics []models.Clinic) []string {
	output := make([]string, len(clinics))

	for index, clinic := range clinics {
		output[index] = clinic.Name
	}

	return output
}
