package main

import "fmt"

func commandExplore(cfg *config, args ...string) error {
	if len(args) != 1 {
		return fmt.Errorf("invalid number of arguments: provide a location name")
	}
	area := args[0]

	location, err := cfg.pokeapiClient.GetLocation(area)
	if err != nil {
		return err
	}

	fmt.Printf("Exploring %v...\n", area)
	fmt.Println("Found Pokemon:")

	for _, pokemonEncounter := range location.PokemonEncounters {
		fmt.Println(" -", pokemonEncounter.Pokemon.Name)
	}

	return nil
}
