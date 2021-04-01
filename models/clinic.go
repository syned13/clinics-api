package models

type ClinicType string

const (
	ClinicTypeDental     ClinicType = "dental"
	ClinicTypeVeterinary ClinicType = "veterinary"
)

type Clinics []Clinic

type Clinic struct {
	Name         string       `json:"name"`
	Statename    string       `json:"stateName"`
	Availability Availability `json:"availability"`
}

type Availability struct {
	From string `json:"from"`
	To   string `json:"to"`
}
