package models

// ClinicType defines the clinic type
type ClinicType string

const (
	// ClinicTypeDental dental
	ClinicTypeDental ClinicType = "dental"
	// ClinicTypeVeterinary veterinary
	ClinicTypeVeterinary ClinicType = "veterinary"
)

// Clinics clinics array
type Clinics []Clinic

// Clinic represents the Clinic entity
type Clinic struct {
	Name         string       `json:"name"`
	Statename    string       `json:"stateName"`
	Availability Availability `json:"availability"`
}

// Availability represents a clinic's availability
type Availability struct {
	From string `json:"from"`
	To   string `json:"to"`
}
