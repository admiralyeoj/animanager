package model

type Media struct {
	Id          string `json:"id"`
	ExternalId  string `json:"id"`
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
