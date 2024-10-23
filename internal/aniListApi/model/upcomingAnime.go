package model

type UpcomingAnime struct {
	Page struct {
		AiringSchedules AiringSchedules `json:"airingSchedules"`
	} `json:"Page"`
}
