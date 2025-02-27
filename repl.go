package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func startRepl() {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Pokedex > ")
		scanner.Scan()
		input := scanner.Text()
		command := cleanInput(input)[0]

		fmt.Printf("Your command was: %s\n", command)
	}
}

func cleanInput(text string) []string {
	cleanInput := strings.Fields(strings.ToLower(text))
	return cleanInput
}
