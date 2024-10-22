package commands

import (
	"github.com/admiralyeoj/anime-announcements/configs"
)

type cliCommand struct {
	name         string
	description  string
	Callback     func(cfg *configs.AniListConfig) error
	CallbackArgs func(cfg *configs.AniListConfig, args ...string) error
}

func GetCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"test": {
			name:        "test",
			description: "Testing command",
			Callback:    Testing,
		},
		"help": {
			name:        "help",
			description: "Displays a help message",
			Callback:    Help,
		},
	}
}
