package model

type UpcomingAnime struct {
	Page struct {
		AiringSchedules []AiringSchedule `json:"airingSchedules"`
	} `json:"Page"`
}
