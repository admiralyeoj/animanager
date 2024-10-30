package command

import (
	"errors"
	"fmt"
	"time"

	"github.com/admiralyeoj/animanager/internal/aniList/service"
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
		Use:   "import-scheduled",
		Short: "Import the upcoming anime",
		Run: func(cmd *cobra.Command, args []string) {
			// Fetching flags using cmd.Flags()
			startDate, _ := cmd.Flags().GetString("start")
			endDate, _ := cmd.Flags().GetString("end")

			if err := c.Handler(startDate, endDate); err != nil {
				fmt.Println(err.Error())
			}
		},
	}

	currentTime := time.Now()
	format := "01/02/2006"
	currentTime.Format(format)

	cmd.Flags().String("start", "", "Start date for importing anime. Default to current date.")
	cmd.Flags().String("end", "", "End date for importing anime. Default to the day after the start date.")

	return cmd
}

// ImportScheduledAnimeHandler handles the scheduled anime import.
func (c *ImportScheduledAnimeCommand) Handler(args ...interface{}) error {

	var startDate, endDate string
	format := "01/02/2006"

	// Check if the start date argument is set
	if len(args) > 0 {
		var ok bool
		startDate, ok = args[0].(string)
		// Check if the first argument is a string
		if !ok {
			return errors.New("Start date set but is not a string.")
		}
	} else {
		startDate = time.Now().Format(format)
	}

	// Check if the first argument is set
	if len(args) > 1 {
		var ok bool
		endDate, ok = args[1].(string)

		if !ok {
			return errors.New("End date set but is not a string.")
		}
	} else {
		date, _ := time.Parse(format, startDate)
		endDate = date.AddDate(0, 0, 1).Format(format)
	}

	err := (*c.aniSrv).ImportUpcomingAnime(startDate, endDate)
	if err != nil {
		return fmt.Errorf("error importing anime: %w", err)
	}

	fmt.Println("Anime successfully imported")
	return nil
}
