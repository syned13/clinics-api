package fetcher

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"

	"github.com/syned13/clinics-api/client"
	"github.com/syned13/clinics-api/models"
)

// DefaultBaseURL default base URl for the API
const DefaultBaseURL = "https://storage.googleapis.com"

var (
	clinicTypePaths = map[models.ClinicType]string{
		models.ClinicTypeVeterinary: "/scratchpay-code-challenge/vet-clinics.json",
		models.ClinicTypeDental:     "/scratchpay-code-challenge/dental-clinics.json",
	}
)

var (
	// ErrInvalidClinicType invalid clinic type
	ErrInvalidClinicType = errors.New("invalid clinic type")
	// ErrCouldNotMakeRequest could not make request
	ErrCouldNotMakeRequest = errors.New("could not make request")
	// ErrCouldNotFetchClinics could not fetch clinics
	ErrCouldNotFetchClinics = errors.New("could not fetch clinics")
	// ErrInvalidStatusCode invalid status code
	ErrInvalidStatusCode = errors.New("invalid status code")
	// ErrCouldNotReadResponse could not read response
	ErrCouldNotReadResponse = errors.New("could not read response")
	// ErrInvalidURL invalid url
	ErrInvalidURL = errors.New("invalid url")
)

// Fetcher represents the entity for fetching clinics
type Fetcher interface {
	FetchClinics(clinicType models.ClinicType) ([]models.Clinic, error)
}

type clinicFetcher struct {
	client  client.Client
	baseURL string
}

// NewFetcher returns a Fecther entity
func NewFetcher(client client.Client, baseURL string) (Fetcher, error) {
	_, err := url.ParseRequestURI(baseURL)
	if err != nil {
		return nil, ErrInvalidURL
	}

	return clinicFetcher{
		client:  client,
		baseURL: baseURL,
	}, nil
}

func (c clinicFetcher) getURL(clinicType models.ClinicType) (string, error) {
	path, ok := clinicTypePaths[clinicType]
	if !ok {
		return "", ErrInvalidClinicType
	}

	url := c.baseURL + path

	return url, nil
}

func (c clinicFetcher) makeRequest(url string) ([]byte, error) {
	request, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return nil, fmt.Errorf("%w: %s", ErrCouldNotMakeRequest, err.Error())
	}

	response, err := c.client.Do(request)
	if err != nil {
		return nil, fmt.Errorf("%w: %s", ErrCouldNotFetchClinics, err.Error())
	}

	defer func() {
		_ = response.Body.Close()
	}()

	if response.StatusCode != http.StatusOK {
		return nil, ErrInvalidStatusCode
	}

	bodyBytes, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, fmt.Errorf("%w: %s", ErrCouldNotReadResponse, err.Error())
	}

	return bodyBytes, nil
}

// FetchClinics fetches the clinics from the storage API
func (c clinicFetcher) FetchClinics(clinicType models.ClinicType) ([]models.Clinic, error) {
	url, err := c.getURL(clinicType)
	if err != nil {
		return nil, err
	}

	body, err := c.makeRequest(url)
	if err != nil {
		return nil, err
	}

	clinics := models.Clinics{}

	err = json.Unmarshal(body, &clinics)
	if err != nil {
		return nil, err
	}

	return clinics, nil
}
