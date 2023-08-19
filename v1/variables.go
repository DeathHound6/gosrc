package v1

type VariableScope struct {
	Type string `json:"type"`
}

type VariableValuesValueFlags struct {
	Miscellaneous bool `json:"miscellaneous"`
}

type VariableValuesValue struct {
	Label string                    `json:"label"`
	Rules string                    `json:"rules"`
	Flags *VariableValuesValueFlags `json:"flags"`
}

type VariableValues struct {
	Values  map[string]*VariableValuesValue `json:"values"`
	Default string                          `json:"default"`
}

type Variable struct {
	ID            string          `json:"id"`
	Name          string          `json:"name"`
	Category      string          `json:"category"`
	Scope         *VariableScope  `json:"scope"`
	Mandatory     bool            `json:"mandatory"`
	UserDefined   bool            `json:"user-defined"`
	Obsoletes     bool            `json:"obsoletes"`
	Values        *VariableValues `json:"values"`
	IsSubcategory bool            `json:"is-subcategory"`
	Links         []*Link         `json:"links"`
}
