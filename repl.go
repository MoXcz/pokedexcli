package main

import (
	"bufio"
	"fmt"
	"github.com/MoXcz/pokedexcli/internal/pokeapi"
	"os"
	"strings"
)

type config struct {
	pokeapiClient pokeapi.Client
	nextURL       *string
	previousURL   *string
}

type cliCommand struct {
	name        string
	description string
	callback    func(*config) error
}

func startRepl(cfg *config) {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Pokedex > ")
		scanner.Scan()
		input := scanner.Text()
		commandName := cleanInput(input)[0]
		command, ok := getValidCommands()[commandName]
		if !ok {
			fmt.Println("Unknown command")
			continue
		}
		err := command.callback(cfg)
		if err != nil {
			fmt.Println("Something went really wrong: ", err)
		}
	}
}

func cleanInput(text string) []string {
	cleanInput := strings.Fields(strings.ToLower(text))
	return cleanInput
}

func getValidCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},
		"help": {
			name:        "help",
			description: "Display a help message",
			callback:    commandHelp,
		},
		"map": {
			name:        "map",
			description: "Display location areas",
			callback:    commandMap,
		},
		"mapb": {
			name:        "mapb",
			description: "Display previous location areas",
			callback:    commandMapb,
		},
	}
}
