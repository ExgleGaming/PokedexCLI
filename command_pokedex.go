package main

import (
	"errors"
	"fmt"
)

func commandPokedex(cfg *config, args ...string) error {
	if len(cfg.caughtPokemon) == 0 {
		return errors.New("You have not caught any pokemon")
	}

	fmt.Println("Your Pokedex:")
	for k, _ := range cfg.caughtPokemon {
		fmt.Printf(" - %v\n", k)
	}
	return nil
}
