package v1

type RunVideosLink struct {
	URI string `json:"uri"`
}

type RunVideos struct {
	Text  string           `json:"text"`
	Links []*RunVideosLink `json:"links"`
}

type RunStatus struct {
	Status     string `json:"status"`
	Examiner   string `json:"examiner"`
	VerifyDate string `json:"verify-date"`
}

type RunPlayers struct {
	Rel  string `json:"rel"`
	Name string `json:"name"`
	ID   string `json:"id"`
	URI  string `json:"uri"`
}

type RunTimes struct {
	Primary          string `json:"primary"`
	PrimaryT         int    `json:"primary_t"`
	Realtime         string `json:"realtime"`
	RealtimeT        int    `json:"realtime_t"`
	RealtimeNoloads  string `json:"realtime_noloads"`
	RealtimeNoloadsT int    `json:"realtime_noloads_t"`
	Ingame           string `json:"ingame"`
	IngameT          int    `json:"ingame_t"`
}

type RunSystem struct {
	Platform string `json:"platform"`
	Emulated bool   `json:"emulated"`
	Region   string `json:"region"`
}

type Run struct {
	ID        string            `json:"id"`
	Weblink   string            `json:"weblink"`
	Game      string            `json:"game"`
	Level     string            `json:"level"`
	Category  string            `json:"category"`
	Videos    *RunVideos        `json:"videos"`
	Comment   string            `json:"comment"`
	Status    *RunStatus        `json:"status"`
	Players   []*RunPlayers     `json:"players"`
	Date      string            `json:"date"`
	Submitted string            `json:"submitted"`
	Times     *RunTimes         `json:"times"`
	System    *RunSystem        `json:"system"`
	Splits    *Link             `json:"splits"`
	Values    map[string]string `json:"values"`
	Links     []*Link           `json:"links"`
}
