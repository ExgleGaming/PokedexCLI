package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

const mapSize = 20

type pokeMap struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

func commandMap() error {
	for i := range mapSize {
		location := getMap(i)
		fmt.Println(location[i].Name)
	}
	return nil
}

func getMap(index int) []pokeMap {
	url := fmt.Sprintf("https://pokeapi.co/api/v2/location-area/%v/", index)
	res, err := http.Get(url)
	if err != nil {
		log.Fatalln(err)
	}
	defer res.Body.Close()

	var currentMap []pokeMap
	decoder := json.NewDecoder(res.Body)
	if err := decoder.Decode(&currentMap); err != nil {
		log.Fatalln(err)
	}
	return currentMap
}
