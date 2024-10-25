package command

import (
	"github.com/admiralyeoj/anime-announcements/internal/repository"
	"github.com/admiralyeoj/anime-announcements/internal/service"
	"github.com/spf13/cobra"
)

type CommandInterface interface {
	Name() string            // Returns the command name
	Command() *cobra.Command // Returns the actual cobra.Command
}

type Command struct {
	Name        string
	Description string
	Handler     func(srvs *service.Services, args []string) error
}

// InitializeCommands sets up all commands and injects repositories into them
func InitializeCommands(repos *repository.Repositories, srvs *service.Services) *cobra.Command {
	// Create the root command
	rootCmd := &cobra.Command{
		Use:   "anime-cli",
		Short: "Manage anime and get upcoming anime. Can be posted to bluesky.",
	}

	// Define all commands implementing CommandInterface
	commands := []CommandInterface{
		NewImportScheduledAnimeCommand(&srvs.AniListSrv),
		NewTestCommand(&repos.BlueSkyRepo),
		// Add other commands implementing CommandInterface here
	}

	// Register each command
	for _, command := range commands {
		rootCmd.AddCommand(command.Command())
	}

	return rootCmd
}
