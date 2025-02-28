package main

import (
	"time"

	"github.com/MatiasSelvaggio/pokedex_REPL/internal/pokeapi"
)

func main() {
	pokeClient := pokeapi.NewClient(5*time.Second, time.Minute*5)
	cfg := &config{
		pokedex:       map[string]Pokemon{},
		pokeapiClient: pokeClient,
	}
	startRepl(cfg)
}
