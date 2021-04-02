package service

import (
	"errors"
	"sync"

	"github.com/syned13/clinics-api/fetcher"
	"github.com/syned13/clinics-api/models"
	"github.com/syned13/clinics-api/repository"
	"github.com/syned13/clinics-api/utils"
)

var (
	// ErrInvalidState invalid state
	ErrInvalidState = errors.New("invalid state")
	// ErrInvalidTime invalid time
	ErrInvalidTime = errors.New("invalid time")
	// ErrMissingTime missing time
	ErrMissingTime = errors.New("missing time")
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

	return s.repo.GetClinics(name, state, from, to)
}

func validateGetClinicsInputs(state string, from, to string) error {
	if state != "" && models.States[state] == "" {
		return ErrInvalidState
	}

	if from != "" && !utils.ValidateHour(from) {
		return ErrInvalidTime
	}

	if to != "" && !utils.ValidateHour(to) {
		return ErrInvalidTime
	}

	return nil
}

func (s clinicService) UpdateClinics(clinics []models.Clinic) error {
	return s.repo.UpdateClinics(clinics)
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
