package v2

type Series struct {
	ID                string  `json:"id"`
	Name              string  `json:"name"`
	URL               string  `json:"url"`
	AddedDate         int     `json:"addedDate"`
	TouchDate         int     `json:"touchDate"`
	WebsiteURL        *string `json:"websiteUrl"`
	DiscordURL        *string `json:"discordUrl"`
	RunCount          int     `json:"runCount"`
	ActivePlayerCount int     `json:"activePlayerCount"`
	TotalPlayerCount  int     `json:"totalPlayerCount"`
	OfficialGameCount int     `json:"officialGameCount"`
	StaticAssets      []Asset `json:"staticAssets"`
}

type SeriesModerator struct {
	SeriesID string             `json:"seriesId"`
	UserID   string             `json:"userId"`
	Level    GameModeratorLevel `json:"level"`
}
