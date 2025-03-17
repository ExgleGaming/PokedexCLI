package main

import (
	"fmt"

	"github.com/exglegaming/PokedexCLI/internal/pokeapi"
)

func commandHelp(client *pokeapi.Client) error {
	fmt.Println()
	fmt.Println("Welcome to the Pokedex!")
	fmt.Println("Usage:")
	fmt.Println()
	for _, cliCmd := range getCommands() {
		fmt.Printf("%s: %s\n", cliCmd.name, cliCmd.description)
	}
	fmt.Println()
	return nil
}
