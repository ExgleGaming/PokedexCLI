package pokeapi

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

type Client struct {
	httpClient *http.Client
	mapSize    int
	baseURL    string
	nextURL    *string
	prevURL    *string
}

func New() *Client {
	return &Client{
		httpClient: &http.Client{
			Timeout: time.Second * 10,
		},
		mapSize: 20,
		baseURL: "https://pokeapi.co/api/v2/",
		nextURL: nil,
		prevURL: nil,
	}
}

type LocationAreasResp struct {
	Count    int     `json:"count"`
	Next     *string `json:"next"`
	Previous *string `json:"previous"`
	Results  []struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"results"`
}

func (c *Client) ListLocationAreas(direction string) ([]string, error) {
	var endpoint string

	// Safeguard against invalid pagination attempts
	if direction == "next" {
		if c.nextURL == nil && c.prevURL == nil {
			// This is the first call
			endpoint = c.baseURL + "location-area"
		} else if c.nextURL == nil {
			return nil, fmt.Errorf("You're already on the last page")
		} else {
			endpoint = *c.nextURL
		}
	} else if direction == "prev" {
		if c.prevURL == nil {
			return nil, fmt.Errorf("You're already on the first page")
		}
		endpoint = *c.prevURL
	} else if direction == "" {
		endpoint = c.baseURL + "location-area"
	}

	res, err := c.httpClient.Get(endpoint)
	if err != nil {
		return nil, fmt.Errorf("Error with GET request to PokeAPI: %v", err)
	}
	defer res.Body.Close()

	// Encoding JSON data to my poke map struct
	var locResp LocationAreasResp
	decoder := json.NewDecoder(res.Body)
	if err := decoder.Decode(&locResp); err != nil {
		return nil, fmt.Errorf("Error with decoding JSON data: %v", err)
	}

	// Now update your client's next/prev URLs
	c.nextURL = locResp.Next
	c.prevURL = locResp.Previous

	// Extract the location names
	locationNames := make([]string, 0, len(locResp.Results))
	for _, result := range locResp.Results {
		locationNames = append(locationNames, result.Name)
	}

	return locationNames, nil
}
