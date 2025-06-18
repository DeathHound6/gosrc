package v2

import (
	"errors"
	"fmt"

	"github.com/DeathHound6/gosrc"
)

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

type GameModerator struct {
	GameID string             `json:"gameId"`
	UserID string             `json:"userId"`
	Level  GameModeratorLevel `json:"level"`
}

func (client *APIClient) GetGameData(params GetGameDataFilters) (*GetGameDataResponse, error) {
	headers := map[string]string{
		"Accept": "application/json",
	}

	filter := map[string]string{}
	if params.GameID != nil {
		filter["gameId"] = *params.GameID
	} else if params.GameURL != nil {
		filter["gameUrl"] = *params.GameURL
	} else {
		return nil, errors.New("either GameID or GameURL must be provided")
	}

	resp, err := gosrc.MakeRequest(gosrc.APIVersionV2, fmt.Sprintf("GetGameData%s", gosrc.MakeURLQuery(filter)), gosrc.HTTPMethodPOST, headers, nil)
	if err != nil {
		return nil, err
	}

	data, err := gosrc.ReadBody[GetGameDataResponse](resp.Body)
	if err != nil {
		return nil, err
	}

	return data, nil
}
