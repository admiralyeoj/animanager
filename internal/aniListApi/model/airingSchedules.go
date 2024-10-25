package model

type AiringSchedules []struct {
	Id              int   `json:"-"`
	ExternalId      int   `json:"id"`
	AiringAt        int   `json:"airingAt"`
	Episode         int   `json:"episode"`
	TimeUntilAiring int   `json:"timeUntilAiring"`
	MediaId         int   `json:"mediaId"`
	Media           Media `json:"media"`
}
