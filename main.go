package main

import (
	"flag"
	"fmt"
	"github.com/xuri/excelize/v2"
	"gopkg.in/yaml.v3"
	"os"
)

func main() {
	PrintVersion()

	flag.Parse()

	fmt.Printf("excel: %v, result: %v, dict: %v, dict_h: %v\n", *e, *r, *d, *h)
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
	if rows, err := f.GetRows("Sheet1"); err == nil {
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
			content := Content{
				data,
				groupName,
				groupId,
				"",
				"",
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
				content.Data = data

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
			contents = append(contents, content)
			model.Content = contents
			model.Coordinate = coordinate
		}
		result, err := yaml.Marshal(model)
		if err = os.WriteFile(*r, result, 0777); err != nil {
			panic(err)
		}
	}
}

func AppendItem(model map[string]map[string][]string, group map[string][]string, attribute []string, row []string) {
	attribute = append(attribute, AppendStr(row[4], row[5], row[6]))
	group[AppendStr(row[2], row[3], "")] = attribute
	model[AppendStr(row[0], row[1], "")] = group
}
