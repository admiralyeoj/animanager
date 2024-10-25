package service

import (
	"fmt"
	"log"
)

func (srv *aniListService) ImportUpcomingAnime(startDate, endDate string) error {
	UpcomingAnimeResp, err := srv.aniListRepository.GetUpcomingAnime(startDate, endDate)

	if err != nil {
		log.Fatal(err)
		return err
	}

	for _, anime := range UpcomingAnimeResp.Page.AiringSchedules {
		media := &anime.Media
		if media.Title.English == "" {
			continue
		}

		err = srv.dbRepositories.Media.UpdateOrCreate(media)
		if err != nil {
			fmt.Println(err.Error())
		}

		err = srv.dbRepositories.AiringSchedule.UpdateOrCreate(media.ID, &anime)
		if err != nil {
			fmt.Println(err.Error())
		}

		err = srv.dbRepositories.MediaTitle.UpdateOrCreate(media.ID, &media.Title)
		if err != nil {
			fmt.Println(err.Error())
		}

		err = srv.dbRepositories.ExternalLinks.UpdateOrCreate(media.ID, &media.ExternalLinks)
		if err != nil {
			fmt.Println(err.Error())
		}
	}

	return nil
}
