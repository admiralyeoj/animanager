package model

// Needed for Anilist api
type UpcomingAnime struct {
	Page struct {
		AiringSchedules []AiringSchedule `json:"airingSchedules"`
	} `json:"Page"`
}
