package service

import (
	"encoding/json"
	"fmt"
	"strconv"
	"time"

	"github.com/admiralyeoj/animanager/internal/database/model"
)

func (srv *blueSkyService) AnnounceAiringAnime() error {

	airing, err := srv.dbRepositories.AiringSchedule.GetNextNotAnnounced()
	if err != nil {
		return err
	}

	// Exit due to no new shows airing
	if airing.ID == 0 {
		return nil
	}

	jsonData, _ := json.Marshal(*airing)
	fmt.Print(string(jsonData))

	// Convert Unix time to time.Time
	t := time.Unix(airing.AiringAt, 0)

	// Format the time.Time object to M/D/Y H:I:s AM/PM
	formattedTime := t.Format("1/2/2006 3:4:05 PM") // M/D/Y H:I:s AM/PM format

	text := airing.Media.Title.English + " Episode " + strconv.Itoa(airing.Episode) + " started airing at " + formattedTime + " EST \n\n"

	var image []string
	if airing.Media.BannerImage != "" {
		image = append(image, airing.Media.BannerImage)
	}

	text += "Streaming at:\n"
	for _, link := range airing.Media.ExternalLinks {
		if link.Type != "STREAMING" {
			continue
		}

		text += "<a href='" + link.Url + "'>" + link.Name + "</a>\n"
	}

	postId, err := srv.bSkyRepository.CreateRecord(&text, &image)
	if err != nil {
		return err
	}

	socialPost := model.SocialPost{
		PostId:    *postId,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
	srv.dbRepositories.SocialPost.Create(airing.ID, &socialPost)

	return nil
}
