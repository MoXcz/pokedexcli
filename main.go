package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println("Hello, World!")
}

func cleanInput(text string) []string {
	cleanInput := strings.Split(strings.TrimSpace(strings.ToLower(text)), " ")
	for i, word := range cleanInput {
		if len(word) == 0 {
			cleanInput = append(cleanInput[:i], cleanInput[i+1:]...)
		}
	}
	return cleanInput
}
