package repository

import (
	"context"
	"log"
	"time"

	aniListModel "github.com/admiralyeoj/anime-announcements/internal/aniListApi/model"
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
						siteUrl
						type
						format
						duration
						episodes
						title {
							english
						}
						coverImage {
							large
						}
						externalLinks {
							site
							url
							type
						}
					}
				
				}
			}
		}
	`)

	req.Var("page", 1)
	req.Var("airingAtGreater", convertDateToTimestamp(startDate))
	req.Var("airingAtLesser", convertDateToTimestamp(endDate))

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

func convertDateToTimestamp(dateStr string) int64 {
	// Define the layout (format) for parsing
	layout := "01/02/2006" // MM/DD/YYYY

	// Parse the date string into a time.Time object
	parsedTime, err := time.Parse(layout, dateStr)
	if err != nil {
		return 0
	}

	// Return the Unix timestamp
	return parsedTime.Unix()
}
