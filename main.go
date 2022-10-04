package main

import (
	"flag"
	"fmt"
	"github.com/xuri/excelize/v2"
	"gopkg.in/yaml.v3"
	"os"
	"strings"
)

func main() {

	flag.Parse()

	fmt.Printf("excel: %v, dict: %v, dict_h: %v\n", *e, *d, *h)

	WriteYaml(ReadExcel())
}

func ReadExcel() map[string]map[string][]string {
	f, err := excelize.OpenFile(*e)
	if err != nil {
		panic(err)
	}
	//model1-group1-attribute[1,2,3]
	//model1-group2-attribute[4,5,6]
	//model2-group1-attribute[7,8,9]
	result := make(map[string]map[string][]string)
	if rows, err := f.GetRows("Template"); err == nil {
		for number, row := range rows {
			if 0 < number {
				if 1 > len(result) || nil == result[AppendStr(row[0], row[1], "")] {
					var attribute []string
					groupList := make(map[string][]string)
					AppendItem(result, groupList, attribute, row)
				} else if nil != result[AppendStr(row[0], row[1], "")] {
					groupList := result[AppendStr(row[0], row[1], "")]
					if nil == groupList[AppendStr(row[2], row[3], "")] {
						var attribute []string
						AppendItem(result, groupList, attribute, row)
					} else if nil != groupList[AppendStr(row[2], row[3], "")] {
						attribute := groupList[AppendStr(row[2], row[3], "")]
						AppendItem(result, groupList, attribute, row)
					} else {
					}
				} else {
				}
			}
		}
	}
	return result
}

func WriteYaml(result map[string]map[string][]string) {
	for mk, mv := range result {
		modelId, modelName, _ := SplitStr(mk)
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
			"",
			"CMDB",
		}
		coordinateX := int8(0)
		coordinateY := int8(0)
		coordinateW := int8(2)
		for gk, gv := range mv {
			groupId, groupName, _ := SplitStr(gk)
			var data []interface{}
			if strings.Contains(groupName, "默认属性") {
				_defaultAttribute := GetDict("default")
				_key := GetTimeStamp()
				_defaultAttribute["key"] = _key
				_defaultAttribute["cruxAttr"] = CruxAttr{
					"名称",
					"ci_name",
					_key,
				}
				data = append(data, _defaultAttribute)
				_defaultCoordinate := Coordinate{}
				_defaultCoordinate.Key = _key
				_defaultCoordinate.X = coordinateX
				coordinateY = coordinateY + GetDictH("default")["h"]
				_defaultCoordinate.Y = coordinateY
				_defaultCoordinate.W = coordinateW
				_defaultCoordinate.H = GetDictH("default")["h"]
				_defaultCoordinate.I = _key
				coordinate = append(coordinate, _defaultCoordinate)
				_cruxAttribute := CruxAttribute{
					true,
					[]KeyWord{
						{
							"名称",
							"ci_name",
							_key,
						},
					},
					false,
				}
				cruxAttributes = append(cruxAttributes, _cruxAttribute)
			}
			_groupKey := GetTimeStamp()
			_groupCoordinate := Coordinate{}
			_groupCoordinate.Key = _groupKey
			_groupCoordinate.X = coordinateX
			coordinateY = coordinateY + GetDictH("group")["h"]
			_groupCoordinate.Y = coordinateY
			_groupCoordinate.W = coordinateW
			_groupCoordinate.H = GetDictH("group")["h"]
			_groupCoordinate.I = _groupKey
			coordinate = append(coordinate, _groupCoordinate)
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
			for _, av := range gv {
				attributeId, attributeName, attributeType := SplitStr(av)
				attribute := GetDict(attributeType)
				attribute["attrID"] = attributeId
				attribute["attrName"] = attributeName
				key := GetTimeStamp()
				attribute["key"] = key
				data = append(data, attribute)

				_coordinate := Coordinate{}
				_coordinate.Key = key
				_coordinate.X = coordinateX
				coordinateY = coordinateY + GetDictH(attributeType)["h"]
				_coordinate.Y = coordinateY
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
		if err = os.WriteFile(modelName+".yaml", result, 0777); err != nil {
			panic(err)
		}
	}
}

func AppendItem(model map[string]map[string][]string, group map[string][]string, attribute []string, row []string) {
	attribute = append(attribute, AppendStr(row[4], row[5], row[6]))
	group[AppendStr(row[2], row[3], "")] = attribute
	model[AppendStr(row[0], row[1], "")] = group
}
