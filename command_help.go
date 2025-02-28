package main

import "fmt"

func commandHelp(cfg *config, args ...string) error {
	usage := ""
	for _, value := range commandRegistry {
		usage += fmt.Sprintf("\n%s: %s", value.name, value.description)
	}
	text := fmt.Sprintf("Welcome to the Pokedex!\nUsage:\n%s", usage)
	fmt.Println(text)
	return nil
}
