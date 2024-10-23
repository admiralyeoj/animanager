package aniListModel

type AiringSchedules []struct {
	ID              int   `json:"id"`
	AiringAt        int   `json:"airingAt"`
	Episode         int   `json:"episode"`
	TimeUntilAiring int   `json:"timeUntilAiring"`
	MediaId         int   `json:"mediaId"`
	Media           Media `json:"media"`
}
