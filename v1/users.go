package v1

type UserNames struct {
	International string `json:"international"`
	Japanese      string `json:"japanese"`
}

type UserNameStyleColor struct {
	Light string `json:"light"`
	Dark  string `json:"dark"`
}

type UserNameStyle struct {
	Style string              `json:"style"`
	Color *UserNameStyleColor `json:"color"`
}

type UserLocationLocation struct {
	Code  string     `json:"code"`
	Names *UserNames `json:"names"`
}

type UserLocation struct {
	Country *UserLocationLocation `json:"country"`
	Region  *UserLocationLocation `json:"region"`
}

type UserConnection struct {
	URI string `json:"uri"`
}

type User struct {
	ID            string          `json:"id"`
	Names         *UserNames      `json:"names"`
	Weblink       string          `json:"weblink"`
	NameStyle     *UserNameStyle  `json:"name-style"`
	Role          string          `json:"role"`
	Signup        string          `json:"signup"`
	Location      *UserLocation   `json:"location"`
	Twitch        *UserConnection `json:"twitch"`
	Hitbox        *UserConnection `json:"hitbox"`
	Youtube       *UserConnection `json:"youtube"`
	Twitter       *UserConnection `json:"twitter"`
	Speedrunslive *UserConnection `json:"speedrunslive"`
	Links         []*Link         `json:"links"`
}
