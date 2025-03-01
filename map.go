package main

import (
	"fmt"
)

func commandMap(cfg *config) error {
	locations, err := cfg.pokeapiClient.GetLocationAreas(cfg.nextURL)
	if err != nil {
		return err
	}

	cfg.nextURL = locations.Next
	cfg.previousURL = locations.Previous

	for _, area := range locations.Results {
		fmt.Println(area.Name)
	}

	return nil
}
func commandMapb(cfg *config) error {
	if cfg.previousURL == nil {
		fmt.Println("you're on the first page")
		return nil
	}
	locations, err := cfg.pokeapiClient.GetLocationAreas(cfg.previousURL)
	if err != nil {
		return err
	}

	cfg.nextURL = locations.Next
	cfg.previousURL = locations.Previous

	for _, area := range locations.Results {
		fmt.Println(area.Name)
	}

	return nil

}
