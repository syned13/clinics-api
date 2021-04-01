package client

import "net/http"

// Client defines the client for making requests
type Client interface {
	Do(req *http.Request) (*http.Response, error)
}
