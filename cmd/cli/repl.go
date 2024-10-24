package main

import (
	"bufio"
	"database/sql"
	"fmt"
	"os"
	"os/signal"
	"strings"
	"syscall"

	"github.com/admiralyeoj/anime-announcements/internal/commands"
	"github.com/admiralyeoj/anime-announcements/internal/config"
	"go.uber.org/zap"
)

func startRepl(log *zap.SugaredLogger, cfg *config.Config, db *sql.DB) {
	reader := bufio.NewScanner(os.Stdin)

	repos, err := config.InitializeRepositories(db)
	if err != nil {
		log.Fatal("Error initializing repositories", zap.Error(err))
	}

	cmds := []commands.Command{
		commands.NewHelpCommand(nil),
		commands.ImportScheduledAnimeCommand,
		// Other commands can be added here
	}

	cmds[0] = commands.NewHelpCommand(cmds)

	// Create a channel to listen for OS interrupts
	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		<-signalChan
		fmt.Println("\nExiting gracefully...")
		os.Exit(0) // Exit without error
	}()

	for {
		fmt.Print(cfg.Name + " > ")
		reader.Scan()

		args := cleanInput(reader.Text())
		if len(args) == 0 {
			continue
		}

		// Parse command line arguments
		commandName := args[0]

		commandFound := false
		for _, cmd := range cmds {
			if cmd.Name == commandName {
				err = cmd.Handler(repos, args[1:]) // Call the handler for the command

				if err != nil {
					log.Fatal(err.Error(), zap.Error(err))
				}

				commandFound = true
				break
			}
		}

		if commandFound {
			continue // Restart the loop after successfully executing a command
		}

		if commandName == "exit" {
			return
		}

		fmt.Printf("Unknown command: %s\n", commandName)

	}
}

func cleanInput(text string) []string {
	output := strings.ToLower(text)
	words := strings.Fields(output)
	return words
}
