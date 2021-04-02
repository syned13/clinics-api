# clinics-api

API for fetching Clinics information

### Prerequisites:
 - Go 1.15 or later
 - Docker (Docker engine min 20.10.2) 
### Running with Docker

`docker build -p {DESIRED_PORT}:5000 -t {DESIRED_TAG} .`

`docker run {DESIRED_TAG}`

### Usage

GET request to localhost:{DESIRED_PORT}/clinics

Available query params:
 - name: clinic name
 - state: state where the clinic is located
 - to: starting opening hour in terms of availabity (in the form of hh:mm)
 - from: ending opening hour in terms of availabity (in the form of hh:mm)

