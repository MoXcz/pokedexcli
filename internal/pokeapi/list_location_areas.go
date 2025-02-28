package pokeapi

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func getLocationAreas(url *string) (RespLocations, error) {
	fullURL := baseURL + "/location-area"
	if url != nil {
		fullURL = *url // in case there's some previous state
	}

	res, err := http.Get(fullURL)
	if err != nil {
		return RespLocations{}, fmt.Errorf("Error communicating with the server: %w", err)
	}
	defer res.Body.Close()

	if res.StatusCode > 299 {
		return RespLocations{}, fmt.Errorf("Response failed with status code: %d and\nbody: %s\n", res.StatusCode, res.Body)
	}

	var areas RespLocations

	decoder := json.NewDecoder(res.Body)
	err = decoder.Decode(&areas)
	if err != nil {
		return RespLocations{}, fmt.Errorf("Error decoding JSON: %w", err)
	}
	return areas, nil
}
