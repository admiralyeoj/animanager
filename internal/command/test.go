package command

import (
	"encoding/json"
	"fmt"

	bSkyRepo "github.com/admiralyeoj/animanager/internal/blueSky/repository"
	bSkySrv "github.com/admiralyeoj/animanager/internal/blueSky/service"
	dbRepos "github.com/admiralyeoj/animanager/internal/database/repository"
	"github.com/spf13/cobra"
)

// TestCommand struct implements CommandInterface
type TestCommand struct {
	dbRepo      *dbRepos.DatabaseRepositories
	blueskyRepo *bSkyRepo.BlueSkyRepository
	blueskySrv  *bSkySrv.BlueSkyService
}

func NewTestCommand(dbRepos dbRepos.DatabaseRepositories, repo *bSkyRepo.BlueSkyRepository, srv *bSkySrv.BlueSkyService) *TestCommand {
	return &TestCommand{
		dbRepo:      &dbRepos,
		blueskyRepo: repo,
		blueskySrv:  srv,
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
			if err := c.Handler(c.dbRepo, c.blueskyRepo, c.blueskySrv); err != nil {
				fmt.Println(err.Error())
			}
		},
	}
}

// ImportScheduledAnimeHandler handles the scheduled anime import.
func (c *TestCommand) Handler(dbRepo *dbRepos.DatabaseRepositories, repo *bSkyRepo.BlueSkyRepository, srv *bSkySrv.BlueSkyService) error {
	results, _ := dbRepo.AiringSchedule.GetNextNotAnnounced()

	jsonObject, _ := json.Marshal(results)

	fmt.Println(string(jsonObject))

	return nil
	images := &[]string{
		"https://s4.anilist.co/file/anilistcdn/media/anime/banner/164172-ceuofxXerReI.jpg",
	}

	_, err := (*repo).CreateRecord("Hello World", images)

	if err != nil {
		fmt.Println(err.Error())
	}

	return nil
}
