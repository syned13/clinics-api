package fetcher

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/syned13/clinics-api/models"
)

var sampleClinics = []models.Clinic{
	{
		Name:      "Sample Clinic",
		Statename: "Sample State",
		Availability: models.Availability{
			From: "12:00",
			To:   "13:00",
		},
	},
}

func TestNewFetcher(t *testing.T) {
	c := require.New(t)

	client := http.Client{}

	_, err := NewFetcher(&client, "someURL")
	c.Equal(ErrInvalidURL, err)

	_, err = NewFetcher(&client, "http://google.com")
	c.Nil(err)
}

func TestFetchClinics(t *testing.T) {
	c := require.New(t)

	server := httptest.NewServer(http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		c.Equal(req.URL.Path, "/scratchpay-code-challenge/dental-clinics.json")

		clinicsB, err := json.Marshal(sampleClinics)
		c.Nil(err)

		rw.Write(clinicsB)
	}))

	defer server.Close()

	f := clinicFetcher{
		client:  server.Client(),
		baseURL: server.URL,
	}

	clinics, err := f.FetchClinics(models.ClinicTypeDental)
	c.Nil(err)
	c.Equal(sampleClinics, clinics)
}

func TestFetchClinicsError(t *testing.T) {
	c := require.New(t)

	server := httptest.NewServer(http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		c.Equal(req.URL.Path, "/scratchpay-code-challenge/dental-clinics.json")

		rw.WriteHeader(http.StatusInternalServerError)
		rw.Write([]byte(""))
	}))

	defer server.Close()

	f := clinicFetcher{
		client:  server.Client(),
		baseURL: server.URL,
	}

	_, err := f.FetchClinics(models.ClinicTypeDental)
	c.Equal(ErrInvalidStatusCode, err)
}
