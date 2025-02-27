package main

import (
	"fmt"
)

func commandMap(cfg *config, args []string) error {
	locationsResp, err := cfg.pokeapiClient.ListLocations(cfg.next)
	if err != nil {
		return err
	}

	cfg.next = locationsResp.Next
	cfg.previous = locationsResp.Previous

	for _, location := range locationsResp.Results {
		fmt.Println(location.Name)
	}
	return nil
}

func commandMapBack(cfg *config, args []string) error {
	locationsResp, err := cfg.pokeapiClient.ListLocations(cfg.previous)
	if err != nil {
		return err
	}

	cfg.next = locationsResp.Next
	cfg.previous = locationsResp.Previous

	for _, location := range locationsResp.Results {
		fmt.Println(location.Name)
	}

	return nil
}
