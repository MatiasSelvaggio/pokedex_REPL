package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/MatiasSelvaggio/pokedex_REPL/internal/pokeapi"
)

const (
	URL = "https://pokeapi.co/api/v2/"
)

var commandRegistry map[string]cliCommand

func startRepl(cfg *config) {

	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Pokedex > ")
		scanner.Scan()
		text := scanner.Text()
		for text == "" {
			scanner.Scan()
			text = scanner.Text()
		}
		splitWords := cleanInput(text)
		actionFound := false
		firstWorld := splitWords[0]
		for key, command := range commandRegistry {
			if key == firstWorld {
				command.callback(cfg, splitWords[1:])
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
			description: "Display pokemons from a location",
			callback:    commandExplore,
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
	callback    func(*config, []string) error
}

type config struct {
	pokeapiClient pokeapi.Client
	next          *string
	previous      *string
}
