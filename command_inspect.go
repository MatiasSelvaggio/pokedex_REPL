package main

import "fmt"

func commandInspect(cfg *config, args ...string) error {
	if len(args) != 1 {
		return fmt.Errorf("you must provide a pokemon name")
	}

	value, ok := cfg.pokedex[args[0]]
	if !ok {
		fmt.Println("you have not caught that pokemon")
		return nil
	}
	fmt.Printf("Name: %s\nHeight: %d\nWeight: %d\nStats:\n", value.Name, value.Height, value.Weight)
	for _, stat := range value.Stats {
		fmt.Printf("  -%s: %d\n", stat.Stat.Name, stat.BaseStat)
	}
	fmt.Println("Types:")
	for _, data := range value.Types {
		fmt.Printf("  - %s\n", data.Type.Name)
	}
	fmt.Printf("You have Capture: %d\n", value.Count)
	return nil

}
