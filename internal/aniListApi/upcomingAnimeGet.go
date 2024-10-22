package aniListApi

import (
	"context"
	"log"
	"time"

	"github.com/machinebox/graphql"
)

type UpcomingAnime struct {
	Page struct {
		AiringSchedules AiringSchedules `json:"airingSchedules"`
	} `json:"Page"`
}

type AiringSchedules []struct {
	ID              int   `json:"id"`
	AiringAt        int   `json:"airingAt"`
	Episode         int   `json:"episode"`
	TimeUntilAiring int   `json:"timeUntilAiring"`
	MediaId         int   `json:"mediaId"`
	Media           Media `json:"media"`
}

type Media struct {
	ID          string `json:"id"`
	SiteUrl     string `json:"siteUrl"`
	Type        string `json:"type"`
	Format      string `json:"format"`
	Duration    int    `json:"duration"`
	Episodes    int    `json:"episodes"`
	BannerImage string `json:"bannerImage"`
	Title       struct {
		English string `json:"english"`
	} `json:"title"`
	CoverImage struct {
		Large string `json:"large"`
	} `json:"coverImage"`
	ExternalLinks []struct {
		Site string `json:"site"`
		Url  string `json:"url"`
		Type string `json:"type"`
	} `json:"externalLinks"`
}

func (c *Client) GetUpcomingAnime(startDate, endDate string) (UpcomingAnime, error) {

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

	var respData UpcomingAnime
	// var respData map[string]interface{}
	if err := c.graphqlClient.Run(ctx, req, &respData); err != nil {
		log.Fatal(err)
	}
	// fmt.Printf("%+v\n", respData)

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
