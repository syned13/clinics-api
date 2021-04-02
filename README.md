# clinics-api

API for fetching Clinics information

### Running
`go run main.go`

### Usage

GET request to localhost:5000/clinics

Available query params:
 - name: clinic name
 - state: state where the clinic is located
 - to: starting opening hour in terms of availabity (in the form of hh:mm)
 - from: ending opening hour in terms of availabity (in the form of hh:mm)

