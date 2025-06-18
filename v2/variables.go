package v2

type Variable struct {
	ID            string                `json:"id"`
	Name          string                `json:"name"`
	URL           string                `json:"url"`
	Pos           int                   `json:"pos"`
	GameID        string                `json:"gameId"`
	Description   *string               `json:"description"`
	CategoryScope VariableCategoryScope `json:"categoryScope"`
	CategoryID    *string               `json:"categoryId"`
	LevelScope    VariableLevelScope    `json:"levelScope"`
	LevelID       *string               `json:"levelId"`
	IsMandatory   bool                  `json:"isMandatory"`
	IsSubcategory bool                  `json:"isSubcategory"`
	IsUserDefined bool                  `json:"isUserDefined"`
	IsObsoleting  bool                  `json:"isObsoleting"`
	DefaultValue  *string               `json:"defaultValue"`
	Archived      bool                  `json:"archived"`
	DisplayMode   VariableDisplayMode   `json:"displayMode"`
}

type Value struct {
	ID         string  `json:"id"`
	Name       string  `json:"name"`
	URL        string  `json:"url"`
	Pos        int     `json:"pos"`
	VariableID string  `json:"variableId"`
	IsMisc     *bool   `json:"isMisc"`
	Rules      *string `json:"rules"`
	Archived   bool    `json:"archived"`
}
