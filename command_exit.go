package main

import (
	"fmt"
	"os"

	"github.com/exglegaming/PokedexCLI/internal/pokeapi"
)

func commandExit(client *pokeapi.Client) error {
	fmt.Print("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}
