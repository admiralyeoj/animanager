package main

type cliCommand struct {
	name         string
	description  string
	Callback     func() error
	CallbackArgs func(args ...string) error
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
