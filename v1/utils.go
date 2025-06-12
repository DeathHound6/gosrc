package v1

import (
	"fmt"

	"github.com/DeathHound6/gosrc"
)

type Link struct {
	Rel string `json:"rel"`
	URI string `json:"uri"`
}

type Asset struct {
	URI    string `json:"uri"`
	Width  int    `json:"width"`
	Height int    `json:"height"`
}

type PaginationResponse struct {
	Offset int     `json:"offset"`
	Max    int     `json:"max"`
	Size   int     `json:"size"`
	Links  []*Link `json:"links"`
}

type Embed string
type Direction string

const (
	EmbedGame       Embed = "game"
	EmbedLevel      Embed = "level"
	EmbedRegion     Embed = "region"
	EmbedCategory   Embed = "category"
	EmbedPlatform   Embed = "platform"
	EmbedCategories Embed = "categories"
	EmbedModerators Embed = "moderators"
	EmbedGametypes  Embed = "gametypes"
	EmbedPlatforms  Embed = "platforms"
	EmbedLevels     Embed = "levels"
	EmbedRegions    Embed = "regions"
	EmbedGenres     Embed = "genres"
	EmbedEngines    Embed = "engines"
	EmbedDevelopers Embed = "developers"
	EmbedPublishers Embed = "publishers"
	EmbedVariables  Embed = "variables"
	EmbedPlayers    Embed = "players"
	// TODO: Add nested embeds

	DirectionAsc  Direction = "asc"
	DirectionDesc Direction = "desc"
)

func ValidateQueryEmbeds(embeds []Embed, allowed []Embed) []string {
	validEmbeds := make([]string, 0)
	for _, embed := range embeds {
		if !gosrc.SliceContains(allowed, embed) {
			fmt.Printf("WARN: Embed '%s' is not allowed for this request\n", string(embed))
			continue
		}
		validEmbeds = append(validEmbeds, string(embed))
	}
	return validEmbeds
}
