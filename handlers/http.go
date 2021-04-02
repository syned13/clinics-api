package handlers

import (
	"fmt"
	"net/http"

	"github.com/syned13/clinics-api/service"
	"github.com/syned13/clinics-api/shared/httputils"
)

type Handler interface {
	HandleGetClinics() http.HandlerFunc
}

type handler struct {
	clinicService service.ClinicService
}

// NewHandler returns a new http handler entity
func NewHandler(serv service.ClinicService) handler {
	return handler{clinicService: serv}
}

// HandleGetClinics handles the get clinics request
func (h handler) HandleGetClinics() http.HandlerFunc {
	return func(rw http.ResponseWriter, r *http.Request) {
		name := r.URL.Query().Get("name")
		state := r.URL.Query().Get("state")
		from := r.URL.Query().Get("from")
		to := r.URL.Query().Get("to")

		clinics, err := h.clinicService.GetClinics(name, state, from, to)
		if err != nil {
			fmt.Println("failed_getting_clinics: " + err.Error())
			httputils.RespondWithError(rw, err)
			return
		}

		httputils.RespondJSON(rw, http.StatusOK, clinics)
	}
}
