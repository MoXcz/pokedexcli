package main

import (
	"fmt"
	"os"
)

func commandExit(cfg *config, args *string) error {
	defer os.Exit(0)
	fmt.Println("Closing the Pokedex... Goodbye!")
	return nil
}
