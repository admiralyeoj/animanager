package commands

import (
	"fmt"
	"strconv"

	"github.com/admiralyeoj/anime-announcements/configs"
)

func Testing(cfg *configs.AniListConfig) error {
	UpcomingAnimeResp, err := cfg.AniListClient.GetUpcomingAnime("10/25/2024", "10/26/2024")

	if err != nil {
		fmt.Println("error")
		return err
	}

	for _, anime := range UpcomingAnimeResp.Page.AiringSchedules {
		fmt.Println("-----------------")
		fmt.Println("ID = " + strconv.Itoa(anime.ID))
		media := anime.Media
		if media.Title.English == "" {
			continue
		}

		fmt.Println(media.SiteUrl)
		fmt.Println("ID = " + strconv.Itoa(anime.ID))
		fmt.Println(media.Title.English)
		fmt.Println(media.Type)
		fmt.Println(media.Format)

		for _, link := range media.ExternalLinks {
			fmt.Println(link.Site + " " + link.Url)
		}

	}

	return nil
}
