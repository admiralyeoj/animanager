package commands

import (
	"fmt"

	"github.com/admiralyeoj/anime-announcements/internal/aniListApi/service"
	"github.com/admiralyeoj/anime-announcements/internal/config"
	"github.com/admiralyeoj/anime-announcements/internal/database/repository"
)

// ImportScheduledAnimeHandler handles the scheduled anime import.
func ImportScheduledAnimeHandler(dbRepo *repository.DatabaseRepositories, args []string) error {
	if len(args) < 2 {
		return fmt.Errorf("missing date arguments")
	}
	startDate := args[0]
	endDate := args[1]

	fmt.Println("test")

	svr := service.NewAniListService(*dbRepo)
	err := svr.ImportUpcomingAnime(startDate, endDate)
	if err != nil {
		return fmt.Errorf("error importing anime: %w", err)
	}
	fmt.Println("Anime successfully imported")
	return nil
}

var ImportScheduledAnimeCommand = Command{
	Name:        "import-anime",
	Description: "Import the upcoming anime",
	Handler: func(repos *config.Repositories, args []string) error {
		return ImportScheduledAnimeHandler(repos.DatabaseRepos, args) // Access using DatabaseRepos
	},
}
