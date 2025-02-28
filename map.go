package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type LocationArea struct {
	Next     string `json:"next"`
	Prev     string `json:"previous"`
	Location []struct {
		Name string `json:"name"`
		Url  string `json:"url"`
	} `json:"results"`
}

func commandMap() error {
	url := "https://pokeapi.co/api/v2/location-area"
	currentState := getValidCommands()["map"].config

	if currentState.Next != "" {
		url = currentState.Next
	}

	res, err := http.Get(url)
	if err != nil {
		return fmt.Errorf("Error communicating with the server: %w", err)
	}
	defer res.Body.Close()

	if res.StatusCode > 299 {
		return fmt.Errorf("Response failed with status code: %d and\nbody: %s\n", res.StatusCode, res.Body)
	}

	var areas LocationArea

	decoder := json.NewDecoder(res.Body)
	err = decoder.Decode(&areas)
	if err != nil {
		return fmt.Errorf("Error decoding JSON: %w", err)
	}

	for _, area := range areas.Location {
		fmt.Println(area.Name)
	}

	currentState.Next = areas.Next
	currentState.Previous = areas.Prev

	return nil
}
func commandMapb() error {
	currentState := getValidCommands()["map"].config
	if currentState.Previous == "" {
		fmt.Println("You're on the first page")
		return nil
	}
	url := currentState.Previous

	res, err := http.Get(url)
	if err != nil {
		return fmt.Errorf("Error communicating with the server: %w", err)
	}
	defer res.Body.Close()

	if res.StatusCode > 299 {
		return fmt.Errorf("Response failed with status code: %d and\nbody: %s\n", res.StatusCode, res.Body)
	}

	var areas LocationArea

	decoder := json.NewDecoder(res.Body)
	err = decoder.Decode(&areas)
	if err != nil {
		return fmt.Errorf("Error decoding JSON: %w", err)
	}

	for _, area := range areas.Location {
		fmt.Println(area.Name)
	}

	currentState.Next = areas.Next
	currentState.Previous = areas.Prev

	return nil
}
