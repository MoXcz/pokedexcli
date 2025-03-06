package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func (c *Client) GetPokemon(pokemonName string) (Pokemon, error) {
	fullURL := baseURL + "/pokemon/" + pokemonName

	body, ok := c.cache.Get(fullURL)

	if !ok {
		req, err := http.NewRequest("GET", fullURL, nil)
		if err != nil {
			return Pokemon{}, err
		}

		res, err := c.httpClient.Do(req)
		if err != nil {
			return Pokemon{}, fmt.Errorf("Error communicating with the server: %w", err)
		}
		defer res.Body.Close()

		body, err = io.ReadAll(res.Body)

		if err != nil {
			return Pokemon{}, fmt.Errorf("Error reading body: %w", err)
		}

		if res.StatusCode > 299 {
			return Pokemon{}, fmt.Errorf("Response failed with status code: %d and\nbody: %s\n", res.StatusCode, body)
		}

		c.cache.Add(fullURL, body)
	}

	var pokemon Pokemon

	err := json.Unmarshal(body, &pokemon)
	if err != nil {
		return Pokemon{}, fmt.Errorf("Error decoding JSON %w", err)
	}

	return pokemon, nil
}
