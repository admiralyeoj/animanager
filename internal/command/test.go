package command

import (
	"fmt"

	"github.com/admiralyeoj/anime-announcements/internal/blueSkyApi/repository"
	"github.com/spf13/cobra"
)

// TestCommand struct implements CommandInterface
type TestCommand struct {
	blueskyRepo *repository.BlueSkyRepository
}

func NewTestCommand(repo *repository.BlueSkyRepository) *TestCommand {
	return &TestCommand{
		blueskyRepo: repo,
	}
}

// Name returns the name of the command
func (c *TestCommand) Name() string {
	return "test"
}

// Command returns the cobra.Command for the command
func (c *TestCommand) Command() *cobra.Command {
	return &cobra.Command{
		Use:   "test",
		Short: "Testing Command",
		Run: func(cmd *cobra.Command, args []string) {
			if err := c.Handler(c.blueskyRepo); err != nil {
				fmt.Println(err.Error())
			}
		},
	}
}

// ImportScheduledAnimeHandler handles the scheduled anime import.
func (c *TestCommand) Handler(repo *repository.BlueSkyRepository) error {
	images := &[]string{
		"https://s4.anilist.co/file/anilistcdn/media/anime/banner/164172-ceuofxXerReI.jpg",
	}

	_, err := (*repo).CreateRecord("Hello World", images)

	if err != nil {
		fmt.Println(err.Error())
	}

	return nil
}
