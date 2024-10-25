package model

type AiringSchedule struct {
	Id         int   `json:"-"`
	ExternalId int   `json:"id"`
	AiringAt   int   `json:"airingAt"`
	Episode    int   `json:"episode"`
	MediaId    uint  `json:"mediaId"`
	Media      Media `json:"media"`
}

func (AiringSchedule) TableName() string {
	return "airing_schedule" // Make sure GORM uses the correct table name
}
