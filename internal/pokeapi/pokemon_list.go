package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func (c *Client) ListPokemons(location string) (ResponsePokemonPokeApi, error) {
	url := baseURL + "/location-area/"
	if location == "" {
		return ResponsePokemonPokeApi{}, fmt.Errorf("error you must send a location")
	}
	url += location
	value, ok := c.cache.Get(url)
	pokemonResp := ResponsePokemonPokeApi{}
	if ok {
		err := json.Unmarshal(value, &pokemonResp)
		if err != nil {
			return ResponsePokemonPokeApi{}, err
		}
		return pokemonResp, nil
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return ResponsePokemonPokeApi{}, err
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return ResponsePokemonPokeApi{}, err
	}
	defer resp.Body.Close()

	dat, err := io.ReadAll(resp.Body)
	if err != nil {
		return ResponsePokemonPokeApi{}, err
	}
	c.cache.Add(url, dat)

	err = json.Unmarshal(dat, &pokemonResp)
	if err != nil {
		return ResponsePokemonPokeApi{}, err
	}

	return pokemonResp, nil
}
