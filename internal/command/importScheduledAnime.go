package command

import (
	"fmt"

	"github.com/admiralyeoj/anime-announcements/internal/aniListApi/service"
	"github.com/spf13/cobra"
)

// ImportScheduledAnimeCommand struct implements CommandInterface
type ImportScheduledAnimeCommand struct {
	aniSrv *service.AniListService
}

func NewImportScheduledAnimeCommand(aniSrv *service.AniListService) *ImportScheduledAnimeCommand {
	return &ImportScheduledAnimeCommand{
		aniSrv: aniSrv,
	}
}

// Name returns the name of the command
func (c *ImportScheduledAnimeCommand) Name() string {
	return "import-anime"
}

// Command returns the cobra.Command for the command
func (c *ImportScheduledAnimeCommand) Command() *cobra.Command {
	return &cobra.Command{
		Use:   "import-anime [start-date] [end-date]",
		Short: "Import the upcoming anime",
		Args:  cobra.MinimumNArgs(2),
		Run: func(cmd *cobra.Command, args []string) {
			if err := importScheduledAnimeHandler(c.aniSrv, args); err != nil {
				fmt.Println(err.Error())
			}
		},
	}
}

// ImportScheduledAnimeHandler handles the scheduled anime import.
func importScheduledAnimeHandler(srv *service.AniListService, args []string) error {
	startDate := args[0]
	endDate := args[1]

	err := (*srv).ImportUpcomingAnime(startDate, endDate)
	if err != nil {
		return fmt.Errorf("error importing anime: %w", err)
	}
	fmt.Println("Anime successfully imported")
	return nil
}
