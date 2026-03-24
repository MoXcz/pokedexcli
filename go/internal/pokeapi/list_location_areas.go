package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func (c *Client) GetLocationAreas(url *string) (RespLocations, error) {
	fullURL := baseURL + "/location-area?offset=0&limit=20"
	if url != nil {
		fullURL = *url // in case there's some previous state
	}

	body, ok := c.cache.Get(fullURL)

	if !ok {
		req, err := http.NewRequest("GET", fullURL, nil)
		if err != nil {
			return RespLocations{}, err
		}

		res, err := c.httpClient.Do(req)
		if err != nil {
			return RespLocations{}, fmt.Errorf("Error communicating with the server: %w", err)
		}
		defer res.Body.Close()

		if res.StatusCode > 299 {
			return RespLocations{}, fmt.Errorf("Response failed with status code: %d and\nbody: %s\n", res.StatusCode, res.Body)
		}

		body, err = io.ReadAll(res.Body)
		if err != nil {
			return RespLocations{}, fmt.Errorf("Error reading body: %w", err)
		}

		c.cache.Add(fullURL, body)
	}

	var locationAreas RespLocations

	err := json.Unmarshal(body, &locationAreas)
	if err != nil {
		return RespLocations{}, fmt.Errorf("Error decoding JSON %w", err)
	}

	return locationAreas, nil
}
