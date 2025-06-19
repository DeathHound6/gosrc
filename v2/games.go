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

type GameBoost struct {
	ID               string   `json:"id"`
	CreatedAt        int      `json:"createdAt"`
	UpdatedAt        int      `json:"updatedAt"`
	GameID           string   `json:"gameId"`
	DonorUserID      string   `json:"donorUserId"`
	Anonymous        bool     `json:"anonymous"`
	RecipientUserIDs []string `json:"recipientUserIds"`
}

type GameStats struct {
	GameID        string `json:"gameId"`
	TotalRuns     int    `json:"totalRuns"`
	TotalRunsFG   int    `json:"totalRunsFG"`
	TotalRunsIL   int    `json:"totalRunsIL"`
	TotalRunTime  int    `json:"totalRunTime"`
	RecentRuns    int    `json:"recentRuns"`
	RecentRunsFG  int    `json:"recentRunsFG"`
	RecentRunsIL  int    `json:"recentRunsIL"`
	TotalPlayers  int    `json:"totalPlayers"`
	ActivePlayers int    `json:"activePlayers"`
	Followers     int    `json:"followers"`
	Guides        int    `json:"guides"`
	Resources     int    `json:"resources"`
}

/*
Get information about a game by its ID or URL path.

Either `GameID` or `GameURL` must be provided.
If both are provided, `GameID` will be used.
*/
func (client *APIClient) GetGameData(params struct{ GameID, GameURL *string }) (*GetGameDataResponse, error) {
	headers := map[string]string{}

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

/*
Get a summary of a game by its ID or URL path.

Either `GameID` or `GameURL` must be provided.
If both are provided, `GameID` will be used.
*/
func (client *APIClient) GetGameSummary(params struct{ GameID, GameURL *string }) (*GetGameSummaryResponse, error) {
	headers := map[string]string{}

	filter := map[string]string{}
	if params.GameID != nil {
		filter["gameId"] = *params.GameID
	} else if params.GameURL != nil {
		filter["gameUrl"] = *params.GameURL
	} else {
		return nil, errors.New("either GameID or GameURL must be provided")
	}

	resp, err := gosrc.MakeRequest(gosrc.APIVersionV2, fmt.Sprintf("GetGameSummary%s", gosrc.MakeURLQuery(filter)), gosrc.HTTPMethodPOST, headers, nil)
	if err != nil {
		return nil, err
	}

	data, err := gosrc.ReadBody[GetGameSummaryResponse](resp.Body)
	if err != nil {
		return nil, err
	}

	return data, nil
}
