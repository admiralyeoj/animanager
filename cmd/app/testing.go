package commands

import (
	"fmt"

	"github.com/admiralyeoj/anime-announcements/internal/aniListApi/service"
)

func Testing() error {
	svr := service.NewAniListService()

	err := svr.ImportUpcomingAnime("10/25/2024", "10/26/2024")

	if err != nil {
		fmt.Println("error")
		return err
	}

	return nil
}
