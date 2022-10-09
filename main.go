package main

import (
	"flag"
	"fmt"
	"github.com/xuri/excelize/v2"
	"gopkg.in/yaml.v3"
	"os"
	"strings"
	"time"
)

func main() {

	flag.Parse()

	fmt.Printf("源Excel: %v, 属性字典: %v, 高度字典: %v\n", *e, *d, *h)

	WriteYaml(ReadExcel())
}

func ReadExcel() map[string]map[string][]string {
	f, err := excelize.OpenFile(*e)
	if err != nil {
		fmt.Println(err.Error())
		panic(err)
	}
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
				AppendItem(result, groupList, attribute, row)
			} else if nil != result[AppendStr(row[0], row[1], "")] {
				//如果模型已存在，获取分组列表
				groupList := result[AppendStr(row[0], row[1], "")]
				if nil == groupList[AppendStr(row[2], row[3], "")] {
					//分组为空，初始化属性列表
					var attribute []string
					AppendItem(result, groupList, attribute, row)
				} else if nil != groupList[AppendStr(row[2], row[3], "")] {
					//分组不为空，添加属性
					attribute := groupList[AppendStr(row[2], row[3], "")]
					AppendItem(result, groupList, attribute, row)
				} else {
					// 不处理
				}
			} else {
				// 不处理
			}
		}
	}
	return result
}

func WriteYaml(result map[string]map[string][]string) {
	_time := time.Now()
	_count := 1
	for mk, mv := range result {
		modelName, modelId, _ := SplitStr(mk)
		var contents []Content
		var coordinate []Coordinate
		var cruxAttributes []CruxAttribute
		model := Model{
			ModelId:        modelId,
			ModelName:      modelName,
			Content:        contents,
			Coordinate:     coordinate,
			CruxAttributes: cruxAttributes,
		}
		coordinateX := int8(0)
		coordinateY := int8(0)
		coordinateW := int8(2)
		for gk, gv := range mv {
			groupName, groupId, _ := SplitStr(gk)
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
				Data:      data,
				GroupName: groupName,
				GroupId:   groupId,
				Key:       _groupKey,
				CruxAttr:  []CruxAttr{},
			}
			for _, av := range gv {
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

		result, err := yaml.Marshal(model)
		if err = os.WriteFile(modelName+".yml", result, 0777); err != nil {
			fmt.Println(err.Error())
			panic(err)
		}
	}
}
