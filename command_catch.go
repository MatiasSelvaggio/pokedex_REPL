package main

import (
	"fmt"
	"math/rand"
)

const (
	chanceToCapture float32 = .2
)

func commandCatch(cfg *config, args ...string) error {
	if len(args) != 1 {
		return fmt.Errorf("you must provide a pokemon name")
	}
	fmt.Printf("Throwing a Pokeball at %s...\n", args[0])
	pokemonResp, err := cfg.pokeapiClient.PokemonData(args[0])
	if err != nil {
		return err
	}

	random := rand.Intn(pokemonResp.BaseExperience)
	base70 := int(chanceToCapture * float32(pokemonResp.BaseExperience))
	if random > base70 {
		fmt.Printf("%s escaped!\n", pokemonResp.Name)
		return nil
	}
	fmt.Printf("%s was captured!\n", pokemonResp.Name)
	value, ok := cfg.pokedex[pokemonResp.Name]
	if !ok {
		cfg.pokedex[pokemonResp.Name] = Pokemon{
			Count:          1,
			Height:         pokemonResp.Height,
			Weight:         pokemonResp.Weight,
			Id:             pokemonResp.Id,
			BaseExperience: pokemonResp.BaseExperience,
			Name:           pokemonResp.Name,
			Stats: []struct {
				BaseStat int
				Effort   int
				Stat     struct {
					Name string
					Url  string
				}
			}(pokemonResp.Stats),
			Types: []struct {
				Slot int
				Type struct {
					Name string
					Url  string
				}
			}(pokemonResp.Types),
		}
	} else {
		value.Count++
		cfg.pokedex[pokemonResp.Name] = value
	}
	// #this display pokedex catch's
	//fmt.Println(pokedex)
	return nil
}
