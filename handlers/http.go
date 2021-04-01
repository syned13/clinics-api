package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/syned13/clinics-api/service"
)

var internalServerError = ErrorResponse{
	Code:    http.StatusInternalServerError,
	Message: "internal server error",
}

// ErrorResponse error response
type ErrorResponse struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

func (e ErrorResponse) Error() string {
	return e.Message
}

// NewBadRequestError returns a bad request error response
func NewBadRequestError(msg string) ErrorResponse {
	return ErrorResponse{
		Code:    http.StatusBadRequest,
		Message: "bad request: " + msg,
	}
}

func NewNotFoundError(resourceName string) ErrorResponse {
	message := "not found"
	if resourceName != "" {
		message = fmt.Sprintf("%s %s", resourceName, message)
	}

	return ErrorResponse{
		Code:    http.StatusNotFound,
		Message: message,
	}
}

// RespondJSON responds with a json with the given status code and data
func RespondJSON(w http.ResponseWriter, statusCode int, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(data)
}

// RespondWithError responds with a json with the given status code and message
func RespondWithError(w http.ResponseWriter, err error) {
	errorResponse, ok := err.(ErrorResponse)
	if !ok {
		RespondInternalServerError(w)
		return
	}

	RespondJSON(w, errorResponse.Code, errorResponse)
}

// RespondInternalServerError responds with an internal server error response
func RespondInternalServerError(w http.ResponseWriter) {
	RespondJSON(w, internalServerError.Code, internalServerError)
}

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

func (h handler) HandleGetClinics() http.HandlerFunc {
	return func(rw http.ResponseWriter, r *http.Request) {
		name := r.URL.Query().Get("name")
		state := r.URL.Query().Get("state")
		from := r.URL.Query().Get("from")
		to := r.URL.Query().Get("to")

		clinics, err := h.clinicService.GetClinics(name, state, from, to)
		if err != nil {
			RespondWithError(rw, err)
			return
		}

		RespondJSON(rw, http.StatusOK, clinics)
	}
}
