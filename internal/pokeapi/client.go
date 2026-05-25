package pokeapi

import (
	"net/http"
	"time"

	"github.com/alexnino8/pokedexcli/internal/pokecache"
)

// client definition
type Client struct {
	httpClient http.Client
	cache      *pokecache.Cache
}

// NewClient initializes a new custom API client
func NewClient(timeout time.Duration) Client {
	return Client{
		httpClient: http.Client{
			Timeout: timeout,
		},
		cache: pokecache.NewCache(5 * time.Second),
	}
}
