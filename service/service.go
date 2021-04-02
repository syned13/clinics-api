package service

import (
	"sync"

	"github.com/syned13/clinics-api/fetcher"
	"github.com/syned13/clinics-api/models"
	"github.com/syned13/clinics-api/repository"
	"github.com/syned13/clinics-api/shared/httputils"
	"github.com/syned13/clinics-api/utils"
)

var (
	// ErrInvalidState invalid state
	ErrInvalidState = httputils.NewBadRequestError("invalid state")
	// ErrInvalidTime invalid time
	ErrInvalidTime = httputils.NewBadRequestError("invalid time")
	// ErrMissingTime missing time
	ErrMissingTime = httputils.NewBadRequestError("missing time")
)

var clinicsToFetch = []models.ClinicType{
	models.ClinicTypeDental,
	models.ClinicTypeVeterinary,
}

type ClinicService interface {
	GetClinics(name string, state string, from, to string) ([]models.Clinic, error)
	UpdateClinics(clinics []models.Clinic) error
	UpdateClinicsFromAPI() error
}

type clinicService struct {
	repo    repository.Repository
	fetcher fetcher.Fetcher
}

// NewClinicService returns a new clinic service entity
func NewClinicService(repo repository.Repository, fetcher fetcher.Fetcher) ClinicService {
	return clinicService{
		repo:    repo,
		fetcher: fetcher,
	}
}

// GetClinics fetches clinics
func (s clinicService) GetClinics(name string, state string, from, to string) ([]models.Clinic, error) {
	err := validateGetClinicsInputs(state, from, to)
	if err != nil {
		return nil, err
	}

	if state != "" && len(state) == 2 {
		state = models.States[state]
	}

	return s.repo.GetClinics(name, state, from, to)
}

func validateGetClinicsInputs(state string, from, to string) error {
	if from != "" && !utils.ValidateHour(from) {
		return ErrInvalidTime
	}

	if to != "" && !utils.ValidateHour(to) {
		return ErrInvalidTime
	}

	return nil
}

func (s clinicService) UpdateClinics(clinics []models.Clinic) error {
	clinics = s.putStatesInLongForm(clinics)

	return s.repo.UpdateClinics(clinics)
}

func (s clinicService) putStatesInLongForm(clinics []models.Clinic) []models.Clinic {
	for index, clinic := range clinics {
		if len(clinic.Statename) > 2 {
			continue
		}

		clinics[index].Statename = models.States[clinic.Statename]
	}

	return clinics
}

func (s clinicService) UpdateClinicsFromAPI() error {
	errChan := make(chan error, len(clinicsToFetch))
	clinicsChan := make(chan []models.Clinic, len(clinicsToFetch))

	wg := sync.WaitGroup{}

	for _, clinicType := range clinicsToFetch {
		wg.Add(1)
		go func(clinicType models.ClinicType) {
			clinics, err := s.fetcher.FetchClinics(clinicType)
			errChan <- err
			clinicsChan <- clinics
			wg.Done()
		}(clinicType)
	}

	wg.Wait()
	close(errChan)
	close(clinicsChan)

	for err := range errChan {
		if err != nil {
			return err
		}
	}

	allClinics := models.Clinics{}

	for clinics := range clinicsChan {
		allClinics = append(allClinics, clinics...)
	}

	err := s.UpdateClinics(allClinics)
	if err != nil {
		return err
	}

	return nil
}
