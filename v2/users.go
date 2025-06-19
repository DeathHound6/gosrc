package v2

import (
	"github.com/DeathHound6/gosrc"
)

type Session struct {
	SignedIn               bool                 `json:"signedIn"`
	ShowAds                bool                 `json:"showAds"`
	User                   *User                `json:"user"`
	Theme                  *Theme               `json:"theme"`
	PowerLevel             SitePowerLevel       `json:"powerLevel"`
	DateFormat             DateFormat           `json:"dateFormat"`
	TimeFormat             TimeFormat           `json:"timeFormat"`
	TimeReference          TimeReference        `json:"timeReference"`
	TimeUnits              TimeDisplayUnits     `json:"timeUnits"`
	HomepageStream         HomepageStreamType   `json:"homepageStream"`
	DisableThemes          bool                 `json:"disableThemes"`
	CSRFToken              string               `json:"csrfToken"`
	NetworkToken           *string              `json:"networkToken"`
	GameList               []Game               `json:"gameList"`
	GameFollowerList       []GameFollower       `json:"gameFollowerList"`
	GameRunnerList         []GameRunner         `json:"gameRunnerList"`
	SeriesList             []Series             `json:"seriesList"`
	SeriesModeratorList    []SeriesModerator    `json:"seriesModeratorList"`
	BoostAvailableTokens   *int                 `json:"boostAvailableTokens"`
	BoostNextTokenDate     int                  `json:"boostNextTokenDate"`
	BoostNextTokenAmount   int                  `json:"boostNextTokenAmount"`
	UserFollowerList       []UserFollower       `json:"userFollowerList"`
	EnabledExperimentIDs   []string             `json:"enabledExperimentIds"`
	ChallengeModeratorList []ChallengeModerator `json:"challengeModeratorList"`
}

type User struct {
	ID           string         `json:"id"`
	Name         string         `json:"name"`
	URL          string         `json:"url"`
	PowerLevel   SitePowerLevel `json:"powerLevel"`
	Pronouns     []string       `json:"pronouns"`
	AreaID       string         `json:"areaId"`
	ColorID1     string         `json:"colorId1"`
	ColorID2     string         `json:"colorId2"`
	IsSupporter  *bool          `json:"isSupporter"`
	ColorAnimate *int           `json:"colorAnimate"`
	IconType     *any           `json:"iconType"`
	OnlineDate   *int           `json:"onlineDate"`
	SignupDate   *int           `json:"signupDate"`
	TouchDate    *int           `json:"touchDate"`
	StaticAssets []Asset        `json:"staticAssets"`
}

type UserFollower struct {
	UserID     string `json:"userId"`
	FollowerID string `json:"followerId"`
}

type Theme struct {
	ID                  string  `json:"id"`
	URL                 string  `json:"url"`
	PrimaryColor        string  `json:"primaryColor"`
	PanelColor          string  `json:"panelColor"`
	PanelOpacity        int     `json:"panelOpacity"`
	NavbarColor         string  `json:"navbarColor"`
	BackgroundColor     string  `json:"backgroundColor"`
	BackgroundFit       int     `json:"backgroundFit"`
	BackgroundPosition  int     `json:"backgroundPosition"`
	BackgroundRepeat    int     `json:"backgroundRepeat"`
	BackgroundScrolling int     `json:"backgroundScrolling"`
	ForegroundFit       int     `json:"foregroundFit"`
	ForegroundPosition  int     `json:"foregroundPosition"`
	ForegroundRepeat    int     `json:"foregroundRepeat"`
	ForegroundScrolling int     `json:"foregroundScrolling"`
	TouchDate           int     `json:"touchDate"`
	StaticAssets        []Asset `json:"staticAssets"`
}

/*
Get the session information for the current user.
*/
func (client *APIClient) GetSession() (*GetSessionResponse, error) {
	headers := map[string]string{}
	client.ApplyAuthorization(headers, nil)

	resp, err := gosrc.MakeRequest(gosrc.APIVersionV2, "GetSession", gosrc.HTTPMethodPOST, headers, nil)
	if err != nil {
		return nil, err
	}

	data, err := gosrc.ReadBody[GetSessionResponse](resp.Body)
	if err != nil {
		return nil, err
	}

	client.csrfToken = &data.Session.CSRFToken

	return data, nil
}
