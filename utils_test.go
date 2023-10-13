package main

import (
	"fmt"
	"testing"
	_ "testing"
	"time"
)

func TestGetDict2(t *testing.T) {
	content := GetDict2("default")
	fmt.Println(content["attrID"])
}

func TestMakeTimeStamp(t *testing.T) {
	stamp := MakeTimeStamp(time.Now())
	fmt.Println(stamp)
}
