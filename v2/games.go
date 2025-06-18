package v2

type Game struct {
	ID                       string          `json:"id"`
	Name                     string          `json:"name"`
	URL                      string          `json:"url"`
	Type                     string          `json:"type"`
	LoadTimes                bool            `json:"loadTimes"`
	Milliseconds             bool            `json:"milliseconds"`
	IGT                      bool            `json:"igt"`
	Verification             bool            `json:"verification"`
	AutoVerify               *bool           `json:"autoVerify"`
	RequireVideo             bool            `json:"requireVideo"`
	Emulator                 EmulatorFilter  `json:"emulator"`
	DefaultTimer             ValidTimer      `json:"defaultTimer"`
	ValidTimers              []ValidTimer    `json:"validTimers"`
	ReleaseDate              *int            `json:"releaseDate"`
	AddedDate                int             `json:"addedDate"`
	TouchDate                int             `json:"touchDate"`
	BaseGameID               *string         `json:"baseGameId"`
	CoverPath                string          `json:"coverPath"`
	Trophy1stPath            *string         `json:"trophy1stPath"`
	Trophy2ndPath            *string         `json:"trophy2ndPath"`
	Trophy3rdPath            *string         `json:"trophy3rdPath"`
	Trophy4thPath            *string         `json:"trophy4thPath"`
	RunCommentsMode          PermissionType  `json:"runCommentsMode"`
	RunCount                 int             `json:"runCount"`
	ActivePlayerCount        int             `json:"activePlayerCount"`
	TotalPlayerCount         int             `json:"totalPlayerCount"`
	BoostReceivedCount       int             `json:"boostReceivedCount"`
	BoostDistinctDonorsCount int             `json:"boostDistinctDonorsCount"`
	Rules                    *string         `json:"rules"`
	ViewPowerLevel           SitePowerLevel  `json:"viewPowerLevel"`
	PlatformIDs              []string        `json:"platformIds"`
	RegionIDs                []string        `json:"regionIds"`
	GameTypeIDs              []GameType      `json:"gameTypeIds"`
	WebsiteURL               *string         `json:"websiteUrl"`
	DiscordURL               *string         `json:"discordUrl"`
	DefaultView              DefaultViewType `json:"defaultView"`
	GuidePermissionType      PermissionType  `json:"guidePermissionType"`
	ResourcePermissionType   PermissionType  `json:"resourcePermissionType"`
	StaticAssets             []Asset         `json:"staticAssets"`
	EmbargoDate              *int            `json:"embargoDate"`
	EmbargoText              *string         `json:"embargoText"`
}

type GameFollower struct {
	GameID         string `json:"gameId"`
	FollowerID     string `json:"followerId"`
	Pos            *int   `json:"pos"`
	AccessCount    int    `json:"accessCount"`
	LastAccessDate int    `json:"lastAccessDate"`
}
