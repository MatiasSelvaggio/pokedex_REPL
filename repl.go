package main

import (
	"fmt"
	"strings"

	"github.com/MatiasSelvaggio/pokedex_REPL/internal/pokeapi"
	"github.com/chzyer/readline"
)

const (
	URL = "https://pokeapi.co/api/v2/"
)

var commandRegistry map[string]cliCommand

func startRepl(cfg *config) {
	scanner, err := readline.New("Pokedex > ")
	if err != nil {
		panic(err)
	}
	defer scanner.Close()

	for {
		text, err := scanner.Readline()
		if err != nil {
			break
		}

		text = strings.TrimSpace(text)
		if text == "" {
			continue
		}

		scanner.SaveHistory(text)

		splitWords := cleanInput(text)
		actionFound := false
		firstWorld := splitWords[0]
		for key, command := range commandRegistry {
			if key == firstWorld {
				command.callback(cfg, splitWords[1:]...)
				actionFound = true
			}
		}
		if !actionFound {
			fmt.Println("Unknown command")
		}

	}
}

func cleanInput(text string) []string {
	lower := strings.ToLower(text)
	words := strings.Fields(lower)
	return words
}

func init() {
	commandRegistry = map[string]cliCommand{
		"help": {
			name:        "help",
			description: "Display a help message",
			callback:    commandHelp,
		},
		"map": {
			name:        "map",
			description: "Display next page of 20 name of maps of pokeapi",
			callback:    commandMap,
		},
		"mapb": {
			name:        "mapb",
			description: "Display previous page of 20 name of maps of pokeapi",
			callback:    commandMapBack,
		},
		"explore": {
			name:        "explore",
			description: "Display a pokemon list from by location",
			callback:    commandExplore,
		},
		"catch": {
			name:        "catch",
			description: "try capturing a pokemon by name or id",
			callback:    commandCatch,
		},
		"inspect": {
			name:        "inspect",
			description: "show data from pokemon you have catch",
			callback:    commandInspect,
		},
		"pokedex": {
			name:        "pokedex",
			description: "show the list of pokemon you had caught",
			callback:    commandPokedex,
		},
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},
	}
}

type cliCommand struct {
	name        string
	description string
	callback    func(*config, ...string) error
}

type config struct {
	pokeapiClient pokeapi.Client
	next          *string
	previous      *string
	pokedex       map[string]Pokemon
}
