package main

import (
	"bufio"
	"fmt"
	"os"
	"os/signal"
	"strings"
	"syscall"

	"github.com/admiralyeoj/animanager/internal/command"
	"github.com/admiralyeoj/animanager/internal/config"
	"github.com/admiralyeoj/animanager/internal/repository"
	"github.com/admiralyeoj/animanager/internal/service"
	"gorm.io/gorm"
)

func startRepl(cfg *config.Config, db *gorm.DB) {

	// Create a root command to serve as the entry point
	repos := repository.InitializeRepositories(db)
	srvs := service.InitializeServices(repos, db)
	rootCmd := command.InitializeCommands(repos, srvs)

	// Create a channel to listen for OS interrupts
	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		<-signalChan
		fmt.Println("\nExiting gracefully...")
		os.Exit(0) // Exit without error
	}()

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
