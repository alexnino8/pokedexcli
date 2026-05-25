package pokeapi

import (
	"net/http"
	"time"
)

// client definition
type Client struct {
	httpClient http.Client
}

// NewClient initializes a new custom API client
func NewClient(timeout time.Duration) Client {
	return Client{
		httpClient: http.Client{
			Timeout: timeout,
		},
	}
}
