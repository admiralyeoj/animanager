package main

import (
	"bufio"
	"database/sql"
	"fmt"
	"os"
	"strings"

	"github.com/admiralyeoj/anime-announcements/config"
	"github.com/admiralyeoj/anime-announcements/internal/command"
	"github.com/admiralyeoj/anime-announcements/internal/repository"
	"github.com/admiralyeoj/anime-announcements/internal/service"
	"go.uber.org/zap"
)

func startRepl(log *zap.SugaredLogger, cfg *config.Config, db *sql.DB) {

	// Create a root command to serve as the entry point
	rootCmd := command.InitializeCommands(service.InitializeServices(repository.InitializeRepositories(db), db))

	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Print(cfg.Name + " > ")
		input, _ := reader.ReadString('\n')

		input = strings.TrimSpace(input)

		if input == "exit" {
			fmt.Println("Exiting...")
			break
		}

		// Split input into arguments
		args := strings.Split(input, " ")

		if len(args) == 0 {
			continue
		}

		rootCmd.SetArgs(args)
		rootCmd.Execute()
	}
}

func cleanInput(text string) []string {
	output := strings.ToLower(text)
	words := strings.Fields(output)
	return words
}
