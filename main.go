package main

import "github.com/exglegaming/PokedexCLI/internal/pokeapi"

func main() {
	pokeClient := pokeapi.New()
	startRepl(pokeClient)
}
