package main

type Model struct {
	ModelId        string          `yaml:"modelId"`
	ModelName      string          `yaml:"modelName"`
	IconPath       string          `default:"icon-1" yaml:"iconPath"`
	Describe       string          `default:"" yaml:"describe"`
	MonitorTypeId  string          `default:"" yaml:"monitorTypeId"`
	Content        []Content       `yaml:"content"`
	Coordinate     []Coordinate    `yaml:"coordinate"`
	CruxAttributes []CruxAttribute `yaml:"cruxAttributes"`
	Category       int8            `default:"0" yaml:"category"`
	AssetType      any             `default:"null" yaml:"assetType"`
	Type           string          `default:"CMDB" yaml:"type"`
}

type Content struct {
	Data         []any      `yaml:"data"`
	GroupName    string     `yaml:"groupName"`
	GroupId      string     `yaml:"groupId"`
	GroupExplain string     `default:"" yaml:"groupExplain"`
	Key          string     `yaml:"key"`
	Group        string     `default:"" yaml:"group"`
	Type         string     `default:"GROUP" yaml:"type"`
	CruxAttr     []CruxAttr `yaml:"cruxAttr"`
}

type Coordinate struct {
	Key    any  `yaml:"key"`
	X      int8 `yaml:"x"`
	Y      int8 `yaml:"y"`
	W      int8 `yaml:"w"`
	H      int8 `yaml:"h"`
	Static bool `yaml:"static"`
	I      any  `yaml:"i"`
}

type CruxAttribute struct {
	MainAttr bool      `default:"true" yaml:"mainAttr"`
	KeyWords []KeyWord `yaml:"keyWords"`
	BuiltIn  bool      `default:"false" yaml:"builtIn"`
}

type CruxAttr struct {
	AttrName string `default:"名称" yaml:"attrName"`
	AttrID   string `default:"ci_name" yaml:"attrID"`
	Key      any    `default:"22222" yaml:"key"`
}

type KeyWord struct {
	AttrName string `yaml:"attrName"`
	AttrID   string `yaml:"attrID"`
	Key      any    `yaml:"key"`
}
