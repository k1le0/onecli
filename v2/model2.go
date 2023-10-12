package v2

type Model2 struct {
	Id               string `json:"id"`
	ModelId          string `json:"modelId"`
	ModelName        string
	IconPath         string
	GroupId          string
	GroupAllName     string
	GroupAllId       string   `json:"groupAllId"`
	AssetType        []string `json:"assetType"`
	ContainsAsset    bool
	Version          int8
	Content          []Content
	UniFieldsGroups  []UniFieldsGroup
	SearchCapability []any
}

type UniFieldsGroup struct {
	Fields        []string
	CanBeEmpty    bool
	SupportImport bool
}

type Content struct {
	AttrID        string
	AttrName      string
	ComponentType string
	ComponentName string
	Explain       string
	AttrInfo      any
	Index         int8
	Properties    []any
}
