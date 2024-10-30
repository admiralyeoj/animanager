package command

import (
	"github.com/admiralyeoj/animanager/internal/repository"
	"github.com/admiralyeoj/animanager/internal/service"
	"github.com/spf13/cobra"
)

type CommandInterface interface {
	Name() string            // Returns the command name
	Command() *cobra.Command // Returns the actual cobra.Command
	Handler(args ...interface{}) error
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
		NewTestCommand(*repos.DatabaseRepos, &repos.BlueSkyRepo, &srvs.BlueSkySrv),
		// Add other commands implementing CommandInterface here
	}

	// Register each command
	for _, command := range commands {
		rootCmd.AddCommand(command.Command())
	}

	return rootCmd
}
