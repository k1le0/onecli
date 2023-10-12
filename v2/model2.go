package v2

type Model2 struct {
	Id               int64 `json:"id"`
	ModelId          string
	ModelName        string
	IconPath         string
	GroupId          string
	GroupAllName     string
	GroupAllId       string
	AssetType        any
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
