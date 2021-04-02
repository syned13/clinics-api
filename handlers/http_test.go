package handlers

import (
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/syned13/clinics-api/models"
	"github.com/syned13/clinics-api/service"
)

func TestHandleGetClinics(t *testing.T) {
	c := require.New(t)

	serv := service.MockService{}

	serv.On("GetClinics", "", "", "", "").Return([]models.Clinic{
		{
			Name:      "Sample Clinic",
			Statename: "CA",
		},
	}, nil)

	handler := NewHandler(&serv)

	req := httptest.NewRequest(http.MethodGet, "/clinics", nil)

	w := httptest.NewRecorder()

	handler.HandleGetClinics()(w, req)
	c.Equal(http.StatusOK, w.Code)
}

func TestHandleGetClinicsError(t *testing.T) {
	c := require.New(t)

	serv := service.MockService{}

	serv.On("GetClinics", "", "", "", "").Return([]models.Clinic{}, errors.New("some error"))

	handler := NewHandler(&serv)
	req := httptest.NewRequest(http.MethodGet, "/clinics", nil)

	w := httptest.NewRecorder()

	handler.HandleGetClinics()(w, req)
	c.Equal(http.StatusInternalServerError, w.Code)
}
