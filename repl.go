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

		commandName := input[0]
		command, ok := getCommands()[commandName]
		if !ok {
			command = getCommands()["help"]
		}

		// Check the command and call it if a usable command
		switch command.name {
		case "exit":
			commandExit()
		case "help":
			commandHelp()
		case "map":
			commandMap()
		default:
			fmt.Println("Unknown command. Please try again")
			continue
		}
	}
}

// This function is to clean up the users input text
func cleanInput(text string) []string {
	lowercased := strings.ToLower(text)
	return strings.Fields(lowercased)
}

// CLI command struct
type cliCommand struct {
	name        string
	description string
	callback    func() error
}

// Function of the CLI commands and returns them. More will be added
func getCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"help": {
			name:        "help",
			description: "Displays a help message",
			callback:    commandHelp,
		},
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},
		"map": {
			name:        "map",
			description: "Displays next 20 map locations",
			callback:    commandMap,
		},
	}
}
