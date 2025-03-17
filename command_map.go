package main

import (
	"fmt"

	"github.com/exglegaming/PokedexCLI/internal/pokeapi"
)

func commandMap(c *pokeapi.Client) error {
	locations, err := c.ListLocationAreas("next")
	if err != nil {
		fmt.Println(err)
		return err
	}

	for _, loc := range locations {
		fmt.Println(loc)
	}
	return nil
}
