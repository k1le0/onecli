package main

import (
	"flag"
	"fmt"
	"gopkg.in/yaml.v3"
	"os"
	"strconv"
	"strings"
	"time"
)

var (
	e  = flag.String("e", "A2机房.xlsx", "target_update.xlsx")
	sf = flag.String("sf", "sourcefile", "source path")
	ef = flag.String("ef", "exportfile", "export path")
	d  = flag.String("d", "dict.yaml", "dict.yaml")
	h  = flag.String("h", "dict_h.yaml", "dict_h.yaml")
)

func MakeTimeStamp(t time.Time) string {
	return strconv.FormatInt(t.UnixNano(), 10)
}

func GetDict(str string) map[string]any {
	file := absD
	yamlFile, err := os.ReadFile(file)
	if err != nil {
		fmt.Println(err.Error())
		panic(err)
	}
	dict := make(map[string]map[string]any)
	if err := yaml.Unmarshal(yamlFile, &dict); err != nil {
		fmt.Println(err.Error())
		panic(err)
	}
	return dict[strings.ToLower(str)]
}

func GetDictH(str string) map[string]int8 {
	file := absH
	yamlFile, err := os.ReadFile(file)
	if err != nil {
		fmt.Println(err.Error())
		panic(err)
	}
	dict := make(map[string]map[string]int8)
	if err := yaml.Unmarshal(yamlFile, &dict); err != nil {
		fmt.Println(err.Error())
		panic(err)
	}
	return dict[strings.ToLower(str)]
}

func AppendStr(str1 string, str2 string, str3 string) string {
	return str1 + "|" + str2 + "|" + str3
}

func SplitStr(str string) (string, string, string) {
	return strings.SplitN(str, "|", 3)[0], strings.SplitN(str, "|", 3)[1], strings.SplitN(str, "|", 3)[2]
}

func AppendItem(model map[string]map[string][]string, group map[string][]string, attribute []string, row []string,
	groupKey map[string][]string, attrKey map[string][]string) {
	attribute = append(attribute, AppendStr(row[4], row[5], row[6]))
	modelIndex := AppendStr(row[0], row[1], "")
	groupIndex := AppendStr(row[2], row[3], "")
	group[groupIndex] = attribute
	model[modelIndex] = group
	if !strings.Contains(strings.Join(groupKey[modelIndex], " "), groupIndex) {
		groupKey[modelIndex] = append(groupKey[modelIndex], groupIndex)
	}
	attrKey[groupIndex] = append(attrKey[groupIndex], AppendStr(row[4], row[5], row[6]))
}

func IsDir(path string) bool {
	s, err := os.Stat(path)
	if err != nil {
		return false
	}
	return s.IsDir()
}

func IsEmptyPath(path string) bool {
	dir, _ := os.ReadDir(path)
	return len(dir) == 0
}

func DirExit(path string) bool {
	_, err := os.Stat(path)
	if err == nil && IsDir(path) {
		return true
	}
	return false
}
