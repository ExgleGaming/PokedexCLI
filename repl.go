package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func startRepl() {
	// Creates a new scanner for user input
	scanner := bufio.NewScanner(os.Stdin)
	for {
		// Start of REPL
		fmt.Print("Pokedex > ")
		scanner.Scan()

		// Cleaning up the input and checking to make sure it is not empty
		input := cleanInput(scanner.Text())
		if len(input) == 0 {
			continue
		}

		// Printing the user's command. This will change soon
		commandName := input[0]
		fmt.Printf("Your command was: %s\n", commandName)
	}
}

// This function is to clean up the users command text
func cleanInput(text string) []string {
	lowercased := strings.ToLower(text)
	return strings.Fields(lowercased)
}
