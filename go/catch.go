package main

import (
	"fmt"
	"math/rand"
)

func commandCatch(cfg *config, args ...string) error {
	if len(args) != 1 {
		return fmt.Errorf("invalid number of arguments: provide a pokemon name")
	}

	pokemonName := args[0]

	pokemon, err := cfg.pokeapiClient.GetPokemon(pokemonName)
	if err != nil {
		return err
	}

	fmt.Printf("Throwing a Pokeball at %s...\n", pokemonName)

	succeess := pokemon.BaseExperience / 50
	chance := rand.Intn(int(succeess) + 1)

	if chance == succeess {
		fmt.Printf("%s was caught\n", pokemonName)
		cfg.caughtPokemon[pokemonName] = pokemon
		fmt.Println("You may now inspect it with the inspect command")
	} else {
		fmt.Printf("%s escpaed\n", pokemonName)
	}

	return nil
}
