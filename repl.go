package main

import (
	"bufio"
	"fmt"
	"github.com/MoXcz/pokedexcli/internal/pokeapi"
	"os"
	"strings"
)

type config struct {
	pokeapiClient   pokeapi.Client
	nextURL         *string
	previousURL     *string
	caughtPokemon   map[string]pokeapi.Pokemon
	historyCommands map[string]string
}

type cliCommand struct {
	name        string
	description string
	callback    func(*config, ...string) error
}

func startRepl(cfg *config) {
	cfg.historyCommands = make(map[string]string)
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Pokedex > ")
		scanner.Scan()

		cmd := cleanInput(scanner.Text())
		if len(cmd) <= 0 {
			continue
		}

		cmdName := cmd[0]
		cmdArgs := []string{}
		if len(cmd) > 1 {
			cmdArgs = cmd[1:]
		}

		validCommand, ok := getValidCommands()[cmdName]
		if !ok {
			fmt.Println("Unknown command")
			continue
		}
		err := validCommand.callback(cfg, cmdArgs...)
		if err != nil {
			fmt.Println(err)
		}
		commandIndex := fmt.Sprintf("%d", len(cfg.historyCommands))
		cfg.historyCommands[commandIndex] = cmdName
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
		"explore": {
			name:        "explore",
			description: "Display area location provided by argument",
			callback:    commandExplore,
		},
		"catch": {
			name:        "catch",
			description: "Try catch action on a pokemon",
			callback:    commandCatch,
		},
		"inspect": {
			name:        "inspect",
			description: "Inspect a caught pokemon",
			callback:    commandInspect,
		},
		"pokedex": {
			name:        "pokedex",
			description: "Print all caught pokemon",
			callback:    commandPokedex,
		},
		"history": {
			name:        "history",
			description: "Print previously used commands",
			callback:    commandHistory,
		},
	}
}
