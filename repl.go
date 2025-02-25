package main

import "strings"

func cleanInput(text string) []string {
	cleanInput := strings.Fields(strings.ToLower(text))
	return cleanInput
}
