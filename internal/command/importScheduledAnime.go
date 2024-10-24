package command

import (
	"fmt"
	"time"

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
	cmd := &cobra.Command{
		Use:   "import-anime [start-date] [end-date]",
		Short: "Import the upcoming anime",
		Run: func(cmd *cobra.Command, args []string) {
			// Fetching flags using cmd.Flags()
			startDate, _ := cmd.Flags().GetString("start")
			endDate, _ := cmd.Flags().GetString("end")

			if err := c.Handler(c.aniSrv, startDate, endDate); err != nil {
				fmt.Println(err.Error())
			}
		},
	}

	currentTime := time.Now()
	format := "01/02/2006"
	currentTime.Format(format)

	cmd.Flags().String("start", "", "Start date for importing anime")
	cmd.Flags().String("end", "", "End date for importing anime")

	return cmd
}

// ImportScheduledAnimeHandler handles the scheduled anime import.
func (c *ImportScheduledAnimeCommand) Handler(srv *service.AniListService, startDate, endDate string) error {

	format := "01/02/2006"

	if startDate == "" {
		startDate = time.Now().Format(format)
	}

	if endDate == "" {
		date, _ := time.Parse(format, startDate)
		endDate = date.AddDate(0, 0, 1).Format(format)
	}

	err := (*srv).ImportUpcomingAnime(startDate, endDate)
	if err != nil {
		return fmt.Errorf("error importing anime: %w", err)
	}

	fmt.Println("Anime successfully imported")
	return nil
}
