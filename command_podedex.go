package main

import "fmt"

func commandPokedex(cfg *config, args ...string) error {
	fmt.Println("You Pokedex:")
	for key, value := range cfg.pokedex {
		fmt.Printf(" - %s owned: %d\n", key, value.Count)
	}
	return nil
}
