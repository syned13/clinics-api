package repository

import (
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/syned13/clinics-api/models"
)

func TestUpdateClinics(t *testing.T) {
	c := require.New(t)

	clinics := []models.Clinic{
		{
			Name:      "Sample Clinic",
			Statename: "Sample State",
			Availability: models.Availability{
				From: "12:00",
				To:   "13:00",
			},
		},
		{
			Name:      "Mayo Clinic",
			Statename: "Sample State",
			Availability: models.Availability{
				From: "12:00",
				To:   "13:00",
			},
		},
		{
			Name:      "Sample Hospital",
			Statename: "Sample State",
			Availability: models.Availability{
				From: "12:00",
				To:   "13:00",
			},
		},
	}

	repo := clinicsRepository{}

	err := repo.UpdateClinics(clinics)
	c.Nil(err)
	c.Equal(len(clinics), len(repo.sortedClinics))
}

func TestGetClinics(t *testing.T) {
	c := require.New(t)

	clinics := []models.Clinic{
		{
			Name:      "Sample Clinic",
			Statename: "IN",
			Availability: models.Availability{
				From: "12:00",
				To:   "13:00",
			},
		},
		{
			Name:      "Mayo Clinic",
			Statename: "CA",
			Availability: models.Availability{
				From: "15:00",
				To:   "13:00",
			},
		},
		{
			Name:      "Sample Hospital",
			Statename: "Sample State",
			Availability: models.Availability{
				From: "12:00",
				To:   "13:00",
			},
		},
	}

	repo := NewClinicRepository()
	repo.UpdateClinics(clinics)

	gottenClinics, _ := repo.GetClinics("clinic", "", "", "")
	c.Len(gottenClinics, 2)

	gottenClinics, _ = repo.GetClinics("", "Sample State", "", "")
	c.Len(gottenClinics, 1)

	gottenClinics, _ = repo.GetClinics("clinic", "IN", "", "")
	c.Len(gottenClinics, 1)

	gottenClinics, _ = repo.GetClinics("clinic", "", "12:00", "12:30")
	c.Len(gottenClinics, 1)
}
