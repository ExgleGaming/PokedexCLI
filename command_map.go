package main

import (
	"fmt"

	"github.com/exglegaming/PokedexCLI/internal/pokeapi"
)

func commandMap(client *pokeapi.Client) error {
	locations, err := client.ListLocationAreas()
	if err != nil {
		return err
	}

	for _, loc := range locations {
		fmt.Println(loc)
	}
	return nil
}
