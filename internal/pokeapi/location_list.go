package pokeapi

import (
	"encoding/json"
	"io"
	"net/http"
)

func (c *Client) ListLocations(pageURL *string) (ResponsePokeApi, error) {
	url := baseURL + "/location-area"
	if pageURL != nil {
		url = *pageURL
	}
	value, ok := c.cache.Get(url)
	locationsResp := ResponsePokeApi{}
	if ok {
		err := json.Unmarshal(value, &locationsResp)
		if err != nil {
			return ResponsePokeApi{}, err
		}
		return locationsResp, nil
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return ResponsePokeApi{}, err
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return ResponsePokeApi{}, err
	}
	defer resp.Body.Close()

	dat, err := io.ReadAll(resp.Body)
	if err != nil {
		return ResponsePokeApi{}, err
	}
	c.cache.Add(url, dat)

	err = json.Unmarshal(dat, &locationsResp)
	if err != nil {
		return ResponsePokeApi{}, err
	}

	return locationsResp, nil
}
