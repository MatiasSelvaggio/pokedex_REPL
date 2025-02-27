package pokeapi

import (
	"net/http"
	"time"

	"github.com/MatiasSelvaggio/pokedex_REPL/internal/pokecache"
)

type Client struct {
	httpClient http.Client
	cache      pokecache.Cache
}

func NewClient(timeout, internal time.Duration) Client {
	return Client{
		httpClient: http.Client{
			Timeout: timeout,
		},
		cache: pokecache.NewCache(internal),
	}
}
