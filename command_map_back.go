package main

import (
	"fmt"

	"github.com/exglegaming/PokedexCLI/internal/pokeapi"
)

func commandMapBack(client *pokeapi.Client) error {
	locations, err := client.ListLocationAreas("prev")
	if err != nil {
		fmt.Println(err)
		return err
	}

	for _, loc := range locations {
		fmt.Println(loc)
	}
	return nil
}
