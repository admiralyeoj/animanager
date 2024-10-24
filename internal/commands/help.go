// internal/commands/help.go
package commands

import (
	"fmt"

	"github.com/admiralyeoj/anime-announcements/internal/config"
)

// HelpHandler handles the help command.
func HelpHandler(cmds []Command) {
	fmt.Println("Available commands:")
	for _, cmd := range cmds {
		fmt.Printf("  %s: %s\n", cmd.Name, cmd.Description)
	}
}

func NewHelpCommand(availableCommands []Command) Command {
	return Command{
		Name:        "help",
		Description: "Display the list of available commands",
		Handler: func(repos *config.Repositories, args []string) error {
			HelpHandler(availableCommands) // Calls HelpHandler with available commands
			return nil
		},
	}
}
