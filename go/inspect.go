package main

import "fmt"

func commandInspect(cfg *config, args ...string) error {
	if len(args) != 1 {
		return fmt.Errorf("invalid number of arguments: provide a pokemon name")
	}

	pokemonName := args[0]

	pokemon, ok := cfg.caughtPokemon[pokemonName]
	if !ok {
		fmt.Println("you have not caught that pokemon")
		return nil
	}

	fmt.Println("Name:", pokemon.Name)
	fmt.Println("Height:", pokemon.Height)
	fmt.Println("Weight:", pokemon.Weight)
	fmt.Println("Stats:")
	for _, stat := range pokemon.Stats {
		fmt.Printf("  - %s: %d\n", stat.Stat.Name, stat.BaseStat)
	}
	fmt.Println("Types:")
	for _, types := range pokemon.Types {
		fmt.Printf("  - %s\n", types.Type.Name)
	}

	return nil
}
