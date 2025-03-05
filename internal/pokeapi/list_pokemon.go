package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func (c *Client) GetPokemon(area string) (RespPokemon, error) {
	fullURL := baseURL + "/location-area/" + area

	body, ok := c.cache.Get(fullURL)

	if !ok {
		req, err := http.NewRequest("GET", fullURL, nil)
		if err != nil {
			return RespPokemon{}, err
		}

		res, err := c.httpClient.Do(req)
		if err != nil {
			return RespPokemon{}, fmt.Errorf("Error communicating with the server: %w", err)
		}
		defer res.Body.Close()

		if res.StatusCode > 299 {
			return RespPokemon{}, fmt.Errorf("Response failed with status code: %d and\nbody: %s\n", res.StatusCode, res.Body)
		}

		body, err = io.ReadAll(res.Body)
		if err != nil {
			return RespPokemon{}, fmt.Errorf("Error reading body: %w", err)
		}

		c.cache.Add(fullURL, body)
	}

	var pokemon RespPokemon

	err := json.Unmarshal(body, &pokemon)
	if err != nil {
		return RespPokemon{}, fmt.Errorf("Error decoding JSON %w", err)
	}

	return pokemon, nil
}
