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

const (
	VERSION = "0.0.1"
)

var (
	e = flag.String("e", "target.xlsx", "target.xlsx")
	r = flag.String("r", "result.yaml", "result.yaml")
	d = flag.String("d", "dict.yaml", "dict.yaml")
	h = flag.String("h", "dict_h.yaml", "dict_h.yaml")
)

func Version() string {
	return VERSION
}

func GetTimeStamp() string {
	_timestamp := time.Now().UnixNano()
	return strconv.FormatInt(_timestamp, 10)
}

func GetDict(str string) map[string]interface{} {
	file := *d
	yamlFile, err := os.ReadFile(file)
	if err != nil {
		panic(err)
	}
	dict := make(map[string]map[string]interface{})
	if err := yaml.Unmarshal(yamlFile, &dict); err != nil {
		panic(err)
	}
	return dict[strings.ToLower(str)]
}

func GetDictH(str string) map[string]int8 {
	file := *h
	yamlFile, err := os.ReadFile(file)
	if err != nil {
		panic(err)
	}
	dict := make(map[string]map[string]int8)
	if err := yaml.Unmarshal(yamlFile, &dict); err != nil {
		panic(err)
	}
	return dict[strings.ToLower(str)]
}

func AppendStr(str1 string, str2 string, str3 string) string {
	return str1 + "_" + str2 + "_" + str3
}

func SplitStr(str string) (string, string, string) {
	return strings.SplitN(str, "_", 3)[0], strings.SplitN(str, "_", 3)[1], strings.SplitN(str, "_", 3)[2]
}

func PrintVersion() {
	fmt.Printf("Onecli %v \n", Version())
}
