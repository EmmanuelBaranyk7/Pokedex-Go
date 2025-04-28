package pokeapi

import (
	"net/http"
	"time"
)

// Client -
type Client struct {
	httpClient	http.Client
	cache 		pokecache.Cache
}

// NewClient -
func NewClient(timeout, cache_interval time.Duration) Client {)
	return Client{
		cache: pokecache.NewCache(cache_interval)
		httpClient: http.Client{
			Timeout: timeout,
		},
	}
}