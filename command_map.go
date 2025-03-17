package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

const mapSize = 20

type pokeMap struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

func commandMap() error {
	for i := range mapSize {
		location, _ := getMap(i)
		fmt.Println(location.Name)
	}
	return nil
}

func getMap(index int) (pokeMap, error) {
	var newPokeMap pokeMap

	// HTTP request to pokeapi to get location area
	url := fmt.Sprintf("https://pokeapi.co/api/v2/location-area/%v/", index)
	res, err := http.Get(url)
	if err != nil {
		return newPokeMap, fmt.Errorf("Error with GET request to PokeAPI: %v", err)
	}
	defer res.Body.Close()

	// Encoding JSON data to my poke map struct
	decoder := json.NewDecoder(res.Body)
	if err := decoder.Decode(&newPokeMap); err != nil {
		return newPokeMap, fmt.Errorf("Error with decoding JSON data: %v", err)
	}

	return newPokeMap, nil
}
