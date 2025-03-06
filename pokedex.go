package main

import "fmt"

func commandPokedex(cfg *config, args ...string) error {
	if len(cfg.caughtPokemon) <= 0 {
		fmt.Println("you have not caught any pokemon")
	}

	for pokemon, _ := range cfg.caughtPokemon {
		println("  -", pokemon)
	}

	return nil
}
