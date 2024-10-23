package service

import (
	"fmt"
	"log"
	"strconv"
)

func (srv *aniListService) ImportUpcomingAnime(startDate, endDate string) error {
	UpcomingAnimeResp, err := srv.aniListRepository.GetUpcomingAnime(startDate, endDate)

	if err != nil {
		log.Fatal(err)
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
