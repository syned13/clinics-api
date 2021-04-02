package fetcher

import (
	"encoding/json"

	"github.com/syned13/clinics-api/models"
)

// Clinic represents clinics methods specific to the fetcher
type Clinic interface {
	ToClinic() models.Clinic
}

// VetClinics veterinary clinics list
type VetClinics []VetClinic

// VetClinic veterinary clinic entity
type VetClinic struct {
	Clinicname string `json:"clinicName"`
	Statecode  string `json:"stateCode"`
	Opening    struct {
		From string `json:"from"`
		To   string `json:"to"`
	} `json:"opening"`
}

// DentalClinics dental clinics list
type DentalClinics []DentalClinic

// DentalClinic dentail clinic entity
type DentalClinic struct {
	Name         string `json:"name"`
	Statename    string `json:"stateName"`
	Availability struct {
		From string `json:"from"`
		To   string `json:"to"`
	} `json:"availability"`
}

func (clinic VetClinic) toClinic() models.Clinic {
	return models.Clinic{
		Name:         clinic.Clinicname,
		Statename:    clinic.Statecode,
		Availability: clinic.Opening,
	}
}

func (clinic DentalClinic) toClinic() models.Clinic {
	return models.Clinic{
		Name:         clinic.Name,
		Statename:    clinic.Statename,
		Availability: clinic.Availability,
	}
}

func unmarshalVetClinics(clinicsB []byte) ([]models.Clinic, error) {
	vetClinics := []VetClinic{}

	err := json.Unmarshal(clinicsB, &vetClinics)
	if err != nil {
		return nil, err
	}

	clinics := make([]models.Clinic, len(vetClinics))

	for index, clinic := range vetClinics {
		clinics[index] = clinic.toClinic()
	}

	return clinics, nil
}

func unmarshalDentalClinics(clinicsB []byte) ([]models.Clinic, error) {
	dentalClinics := []DentalClinic{}

	err := json.Unmarshal(clinicsB, &dentalClinics)
	if err != nil {
		return nil, err
	}

	clinics := make([]models.Clinic, len(dentalClinics))

	for index, clinic := range dentalClinics {
		clinics[index] = clinic.toClinic()
	}

	return clinics, nil
}

func unmarshalListOfClinics(clinicsB []byte, clinicType models.ClinicType) ([]models.Clinic, error) {
	switch clinicType {
	case models.ClinicTypeDental:
		return unmarshalDentalClinics(clinicsB)
	case models.ClinicTypeVeterinary:
		return unmarshalVetClinics(clinicsB)
	}

	return nil, ErrInvalidClinicType
}
