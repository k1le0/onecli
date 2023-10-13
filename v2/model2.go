package v2

type Model2 struct {
	Id               string           `json:"id"`
	ModelId          string           `json:"modelId"`
	ModelName        string           `json:"modelName"`
	IconPath         string           `json:"iconPath"`
	GroupId          string           `json:"groupId"`
	GroupAllName     string           `json:"groupAllName"`
	GroupAllId       string           `json:"groupAllId"`
	AssetType        []string         `json:"assetType"`
	ContainsAsset    bool             `json:"containsAsset"`
	Version          int8             `json:"version"`
	Content          []Content        `json:"content"`
	UniFieldsGroups  []UniFieldsGroup `json:"uniFieldsGroups"`
	SearchCapability []any            `json:"searchCapability"`
}

type UniFieldsGroup struct {
	Fields        []string `json:"fields"`
	CanBeEmpty    bool     `json:"canBeEmpty"`
	SupportImport bool     `json:"supportImport"`
}

type Content struct {
	AttrID        string `json:"attrID"`
	AttrName      string `json:"attrName"`
	ComponentType string `json:"componentType"`
	ComponentName string `json:"componentName"`
	Explain       string `json:"explain"`
	AttrInfo      any    `json:"attrInfo"`
	Index         int8   `json:"index"`
	Properties    []any  `json:"properties"`
}

type AttrInfo struct {
	DefaultValue     string     `json:"defaultValue"`
	Placeholder      string     `json:"placeholder"`
	FieldWidth       string     `json:"fieldWidth"`
	OptionType       []string   `json:"optionType"`
	CheckType        []string   `json:"checkType"`
	ConditionContent any        `json:"conditionContent"`
	ValidateRuleId   string     `json:"validateRuleId"`
	ValidateRule     string     `json:"validateRule"`
	Range            []int8     `json:"range"`
	DictionaryName   string     `json:"dictionaryName"`
	DictionaryId     string     `json:"dictionaryId"`
	EnumList         []EnumItem `json:"enumList"`
}

type EnumItem struct {
	DictId     string `json:"dictId"`
	Level      int8   `json:"level"`
	ItemCode   string `json:"itemCode"`
	ItemName   string `json:"itemName"`
	ItemColour string `json:"itemColour"`
	Children   any    `json:"children"`
}
