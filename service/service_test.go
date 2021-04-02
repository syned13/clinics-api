package service

import (
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/syned13/clinics-api/fetcher"
	"github.com/syned13/clinics-api/models"
	"github.com/syned13/clinics-api/repository"
)

func TestGetClinics(t *testing.T) {
	c := require.New(t)

	repo := repository.MockRepository{}
	fetcher := fetcher.MockFetcher{}

	expectedClinics := []models.Clinic{
		{
			Name:      "Monica",
			Statename: "California",
		},
	}

	repo.On("GetClinics", "Monica", "", "", "").Return(expectedClinics, nil)

	service := NewClinicService(&repo, &fetcher)
	clinics, err := service.GetClinics("Monica", "", "", "")

	c.Nil(err)
	c.Equal(expectedClinics, clinics)
}
