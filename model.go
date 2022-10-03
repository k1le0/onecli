package main

type Model struct {
	ModelId        string          `yaml:"modelId"`
	ModelName      string          `yaml:"modelName"`
	IconPath       string          `yaml:"iconPath"`
	Describe       string          `yaml:"describe"`
	MonitorTypeId  string          `yaml:"monitorTypeId"`
	Content        []Content       `yaml:"content"`
	Coordinate     []Coordinate    `yaml:"coordinate"`
	CruxAttributes []CruxAttribute `yaml:"cruxAttributes"`
	Category       int8            `yaml:"category"`
	AssetType      string          `yaml:"assetType"`
	Type           string          `yaml:"type"`
}

type Content struct {
	Data         []interface{} `yaml:"data"`
	GroupName    string        `yaml:"groupName"`
	GroupId      string        `yaml:"groupId"`
	GroupExplain string        `yaml:"groupExplain"`
	Key          string        `yaml:"key"`
	Group        string        `yaml:"group"`
	Type         string        `yaml:"type"`
	CruxAttr     []CruxAttr    `yaml:"cruxAttr"`
}

type Coordinate struct {
	Key    string `yaml:"key"`
	X      int8   `yaml:"x"`
	Y      int8   `yaml:"y"`
	W      int8   `yaml:"w"`
	H      int8   `yaml:"h"`
	Static bool   `yaml:"static"`
	I      string `yaml:"i"`
}

type CruxAttribute struct {
	MainAttr bool      `yaml:"mainAttr"`
	KeyWords []KeyWord `yaml:"keyWords"`
	BuiltIn  bool      `yaml:"builtIn"`
}

type CruxAttr struct {
	AttrName string `yaml:"attrName"`
	AttrID   string `yaml:"attrID"`
	Key      string `yaml:"key"`
}

type KeyWord struct {
	AttrName string `yaml:"attrName"`
	AttrID   string `yaml:"attrID"`
	Key      string `yaml:"key"`
}
