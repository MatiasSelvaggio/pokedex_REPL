package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func (c *Client) PokemonData(name string) (ResponsePokemonDataApi, error) {
	if name == "" {
		return ResponsePokemonDataApi{}, fmt.Errorf("you must send a name")
	}
	url := baseURL + "/pokemon/" + name
	value, ok := c.cache.Get(url)
	pokemonResp := ResponsePokemonDataApi{}
	if ok {
		err := json.Unmarshal(value, &pokemonResp)
		if err != nil {
			return ResponsePokemonDataApi{}, err
		}
		return pokemonResp, nil
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return ResponsePokemonDataApi{}, err
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return ResponsePokemonDataApi{}, err
	}

	defer resp.Body.Close()

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return ResponsePokemonDataApi{}, err
	}
	c.cache.Add(url, data)

	err = json.Unmarshal(data, &pokemonResp)
	if err != nil {
		return ResponsePokemonDataApi{}, err
	}
	return pokemonResp, nil
}
