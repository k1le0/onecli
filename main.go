package main

import (
	"flag"
	"fmt"
	v2 "github.com/k1le0/onecli/v2"
	"github.com/xuri/excelize/v2"
	"gopkg.in/yaml.v3"
	"os"
	"path/filepath"
	"strings"
	"time"
)

func main() {

	flag.Parse()

	absE, _ = filepath.Abs(*e)

	absSf, _ = filepath.Abs(*sf)

	absEf, _ = filepath.Abs(*ef)

	absD, _ = filepath.Abs(*d)

	absH, _ = filepath.Abs(*h)

	fmt.Printf("源文件夹: %v, 目标文件夹: %v, 源Excel: %v, 属性字典: %v, 高度字典: %v\n", absSf, absEf, absE, absD, absH)

	DealData()
}

var absE string
var absSf string
var absEf string
var absD string
var absH string

func DealData() {
	if IsDir(absSf) && !IsEmptyPath(absSf) {
		dir, _ := os.ReadDir(absSf)
		for _, item := range dir {
			WriteYaml(ReadExcelFile(filepath.Join(absSf, item.Name())))
		}
	} else {
		WriteYaml(ReadExcelFile(absE))
	}
}

func ReadExcelFile(file string) (map[string][]string, map[string][]string, map[string]map[string][]string) {
	f, err := excelize.OpenFile(file)
	if err != nil {
		fmt.Println(err.Error())
		panic(err)
	}
	groupKey := make(map[string][]string)
	attrKey := make(map[string][]string)
	result := make(map[string]map[string][]string)
	rows, err := f.GetRows("Template")
	if err != nil {
		fmt.Println(err.Error())
		panic(err)
	}
	for number, row := range rows {
		if 0 < number {
			if 1 > len(result) || nil == result[AppendStr(row[0], row[1], "")] {
				//如果模型不存在，初始化分组列表
				var attribute []string
				groupList := make(map[string][]string)
				AppendItem(result, groupList, attribute, row, groupKey, attrKey)
			} else if nil != result[AppendStr(row[0], row[1], "")] {
				//如果模型已存在，获取分组列表
				groupList := result[AppendStr(row[0], row[1], "")]
				if nil == groupList[AppendStr(row[2], row[3], "")] {
					//分组为空，初始化属性列表
					var attribute []string
					AppendItem(result, groupList, attribute, row, groupKey, attrKey)
				} else if nil != groupList[AppendStr(row[2], row[3], "")] {
					//分组不为空，添加属性
					attribute := groupList[AppendStr(row[2], row[3], "")]
					AppendItem(result, groupList, attribute, row, groupKey, attrKey)
				} else {
					// 不处理
				}
			} else {
				// 不处理
			}
		}
	}
	return groupKey, attrKey, result
}

func WriteYaml(groupKey map[string][]string, attrKey map[string][]string, result map[string]map[string][]string) {
	_time := time.Now()
	_count := 1
	for mk, _ := range result {
		modelName, modelId, _ := SplitStr(mk)
		var contents []Content
		var coordinate []Coordinate
		var cruxAttributes []CruxAttribute
		model := Model{
			modelId,
			modelName,
			"icon-1",
			"",
			"",
			contents,
			coordinate,
			cruxAttributes,
			0,
			nil,
			"CMDB",
		}
		coordinateX := int8(0)
		coordinateY := int8(0)
		coordinateW := int8(2)
		for _, value := range groupKey[mk] {
			groupName, groupId, _ := SplitStr(value)
			var data []any
			_groupKey := MakeTimeStamp(_time.Add(time.Duration(_count) * time.Minute))
			_count++
			if strings.Contains(groupName, "默认属性") {
				_groupKey = "11111"
			}
			_groupCoordinate := Coordinate{}
			_groupCoordinate.Key = _groupKey
			_groupCoordinate.X = coordinateX
			_groupCoordinate.Y = coordinateY
			coordinateY = coordinateY + GetDictH("group")["h"]
			_groupCoordinate.W = coordinateW
			_groupCoordinate.H = GetDictH("group")["h"]
			_groupCoordinate.I = _groupKey
			coordinate = append(coordinate, _groupCoordinate)
			if strings.Contains(groupName, "默认属性") {
				_defaultAttribute := GetDict("default")
				_key := _defaultAttribute["key"]
				data = append(data, _defaultAttribute)
				_defaultCoordinate := Coordinate{}
				_defaultCoordinate.Key = _key
				_defaultCoordinate.X = coordinateX
				_defaultCoordinate.Y = coordinateY
				coordinateY = coordinateY + GetDictH("group")["h"]
				_defaultCoordinate.W = coordinateW
				_defaultCoordinate.H = GetDictH("default")["h"]
				_defaultCoordinate.I = _key
				coordinate = append(coordinate, _defaultCoordinate)
				_cruxAttribute := CruxAttribute{
					KeyWords: []KeyWord{
						{
							"名称",
							"ci_name",
							_key,
						},
					},
				}
				cruxAttributes = append(cruxAttributes, _cruxAttribute)
			}
			content := Content{
				data,
				groupName,
				groupId,
				"",
				_groupKey,
				"",
				"GROUP",
				[]CruxAttr{},
			}
			for _, av := range attrKey[value] {
				attributeName, attributeId, attributeType := SplitStr(av)
				attribute := GetDict(attributeType)
				attribute["attrID"] = attributeId
				attribute["attrName"] = attributeName
				key := MakeTimeStamp(_time.Add(time.Duration(_count) * time.Minute))
				_count++
				attribute["key"] = key
				data = append(data, attribute)

				_coordinate := Coordinate{}
				_coordinate.Key = key
				_coordinate.X = coordinateX
				_coordinate.Y = coordinateY
				coordinateY = coordinateY + GetDictH(attributeType)["h"]
				_coordinate.W = coordinateW
				_coordinate.H = GetDictH(attributeType)["h"]
				_coordinate.I = key
				coordinate = append(coordinate, _coordinate)
			}
			content.Data = data
			contents = append(contents, content)
		}
		model.Content = contents
		model.Coordinate = coordinate
		model.CruxAttributes = cruxAttributes

		if !DirExit(absEf) {
			err := os.Mkdir(absEf, os.ModePerm)
			if err != nil {
				fmt.Println(err.Error())
				panic(err)
			}
		}

		result, err := yaml.Marshal(model)
		if err = os.WriteFile(filepath.Join(absEf, modelName+".yml"), result, 0777); err != nil {
			fmt.Println(err.Error())
			panic(err)
		}
	}
}

func WriteYaml2(groupKey map[string][]string, attrKey map[string][]string, result map[string]map[string][]string) {
	_time := time.Now()
	for mk, _ := range result {
		modelName, modelId, _ := SplitStr(mk)
		var contents []v2.Content
		var uniFieldsGroups []v2.UniFieldsGroup
		model := v2.Model2{
			Id:               MakeTimeStamp(_time),
			ModelId:          modelId,
			ModelName:        modelName,
			IconPath:         "icon-1",
			GroupId:          "",
			GroupAllName:     "",
			GroupAllId:       "",
			AssetType:        nil,
			ContainsAsset:    false,
			Version:          13,
			Content:          contents,
			UniFieldsGroups:  nil,
			SearchCapability: nil,
		}
		for _, value := range groupKey[mk] {
			groupName, groupId, _ := SplitStr(value)
			var properties []any
			content := v2.Content{
				AttrID:        groupId,
				AttrName:      groupName,
				ComponentType: "ATTRIBUTE_GROUP",
				ComponentName: "属性分组",
				Explain:       "",
				AttrInfo:      nil,
				Index:         0,
				Properties:    properties,
			}
			// 基本信息 属性分组默认包含一个名称和文本
			if strings.Contains(groupName, "基本信息") {

			}
			for _, av := range attrKey[value] {
				attributeName, attributeId, attributeType := SplitStr(av)
				attribute := GetDict(attributeType)
				attribute["attrID"] = attributeId
				attribute["attrName"] = attributeName
				properties = append(properties, attribute)
			}
			content.Properties = properties

		}
		model.Content = contents
		model.UniFieldsGroups = uniFieldsGroups
	}
}
