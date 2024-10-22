package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/admiralyeoj/anime-announcements/commands"
	"github.com/admiralyeoj/anime-announcements/configs"
)

// cliName is the name used in the repl prompts
const cliName string = "Anime Announcements"

func startRepl(cfg *configs.AniListConfig) {
	reader := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print(cliName + " > ")
		reader.Scan()

		words := cleanInput(reader.Text())
		if len(words) == 0 {
			continue
		}

		commandName := words[0]

		command, exists := commands.GetCommands()[commandName]
		if exists {
			if command.Callback != nil {
				err := command.Callback(cfg)
				if err != nil {
					fmt.Println(err)
				}
			} else if command.CallbackArgs != nil {
				args := []string{}
				if len(words) > 1 {
					args = words[1:]
				}

				err := command.CallbackArgs(cfg, args...)
				if err != nil {
					fmt.Println(err)
				}
			}

			continue
		} else {
			fmt.Println(commandName, ": command not found")
			continue
		}
	}
}

func cleanInput(text string) []string {
	output := strings.ToLower(text)
	words := strings.Fields(output)
	return words
}
