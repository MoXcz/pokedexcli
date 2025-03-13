package main

import "fmt"

func commandHistory(cfg *config, args ...string) error {
	for i, command := range cfg.historyCommands {
		fmt.Println(i, command)

	}
	return nil
}
