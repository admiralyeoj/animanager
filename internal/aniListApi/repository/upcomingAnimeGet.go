package repository

import (
	"context"
	"log"
	"time"

	aniListModel "github.com/admiralyeoj/anime-announcements/internal/aniListApi/model"
	"github.com/admiralyeoj/anime-announcements/util"
	"github.com/machinebox/graphql"
)

func (repo *aniListRepository) GetUpcomingAnime(startDate, endDate string) (aniListModel.UpcomingAnime, error) {

	// make a request
	req := graphql.NewRequest(`
		query Query($page: Int, $airingAtGreater: Int, $airingAtLesser: Int){
			Page(page: $page) {
				airingSchedules(airingAt_greater: $airingAtGreater, airingAt_lesser: $airingAtLesser) {
					id
					airingAt
					episode
					timeUntilAiring
					media {
						id
						siteUrl
						type
						format
						duration
						episodes
						bannerImage
						title {
							english
						}
						coverImage {
							large
						}
						externalLinks {
							siteId
							site
							url
							type
							language
						}
					}
				
				}
			}
		}
	`)

	req.Var("page", 1)
	req.Var("airingAtGreater", util.ConvertDateToTimestamp(startDate))
	req.Var("airingAtLesser", util.ConvertDateToTimestamp(endDate))

	req.Header.Set("Cache-Control", "no-cache")

	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()

	var respData aniListModel.UpcomingAnime
	// var respData map[string]interface{}
	if err := repo.graphqlClient.Run(ctx, req, &respData); err != nil {
		log.Fatal(err)
	}

	return respData, nil
}
