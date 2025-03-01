package pokeapi

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func (c *Client) getLocationAreas(url *string) (RespLocations, error) {
	fullURL := baseURL + "/location-area"
	if url != nil {
		fullURL = *url // in case there's some previous state
	}

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

	var locationAreas RespLocations
	decoder := json.NewDecoder(res.Body)
	err = decoder.Decode(&locationAreas)
	if err != nil {
		return RespLocations{}, fmt.Errorf("Error decoding JSON: %w", err)
	}
	return locationAreas, nil
}
