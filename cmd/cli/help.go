package main

import (
	"fmt"
)

func Help() error {
	fmt.Println()
	fmt.Println("Welcome to the Anime Announcer!")
	fmt.Println("Usage:")
	fmt.Println()

	for _, cmd := range GetCommands() {
		fmt.Printf("%s: %s\n", cmd.name, cmd.description)
	}

	fmt.Println()
	return nil
}
