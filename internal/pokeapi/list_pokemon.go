package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func (c *Client) GetLocation(area string) (RespLocation, error) {
	fullURL := baseURL + "/location-area/" + area

	body, ok := c.cache.Get(fullURL)

	if !ok {
		req, err := http.NewRequest("GET", fullURL, nil)
		if err != nil {
			return RespLocation{}, err
		}

		res, err := c.httpClient.Do(req)
		if err != nil {
			return RespLocation{}, fmt.Errorf("Error communicating with the server: %w", err)
		}
		defer res.Body.Close()

		if res.StatusCode > 299 {
			return RespLocation{}, fmt.Errorf("Response failed with status code: %d and\nbody: %s\n", res.StatusCode, res.Body)
		}

		body, err = io.ReadAll(res.Body)
		if err != nil {
			return RespLocation{}, fmt.Errorf("Error reading body: %w", err)
		}

		c.cache.Add(fullURL, body)
	}

	var pokemon RespLocation

	err := json.Unmarshal(body, &pokemon)
	if err != nil {
		return RespLocation{}, fmt.Errorf("Error decoding JSON %w", err)
	}

	return pokemon, nil
}
