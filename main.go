package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/subosito/gotenv"
	"github.com/syned13/clinics-api/fetcher"
	"github.com/syned13/clinics-api/handlers"
	"github.com/syned13/clinics-api/repository"
	"github.com/syned13/clinics-api/service"
)

var port string

const defaultPort = "5000"

func init() {
	gotenv.Load()

	port = os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}
}

func main() {

	router := mux.NewRouter()

	repo := repository.NewClinicRepository()

	fetcher, err := fetcher.NewFetcher(http.DefaultClient, fetcher.DefaultBaseURL)
	if err != nil {
		log.Fatal(err)
	}

	service := service.NewClinicService(repo, fetcher)
	handler := handlers.NewHandler(service)

	router.Handle("/clinics", handler.HandleGetClinics())

	err = service.UpdateClinicsFromAPI()
	if err != nil {
		log.Fatal(err)
	}

	l := log.Default()
	l.Print("Done updating...")

	// gocron.Every(1).Day().At("00:00").Do(service.)
	err = http.ListenAndServe(fmt.Sprintf(":%s", port), router)
	if err != nil {
		log.Fatal(err)
	}
}
