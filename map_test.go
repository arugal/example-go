package example_go

import (
	"fmt"
	"testing"
)

func TestMap1(t *testing.T) {

	var countryCapitalMap map[string]string
	countryCapitalMap = make(map[string]string)

	countryCapitalMap["France"] = "巴黎"
	countryCapitalMap["Italy"] = "罗马"
	countryCapitalMap["Japan"] = "东京"
	countryCapitalMap["India "] = "新德里"

	for k, v := range countryCapitalMap {
		fmt.Println("range1:", k, "-", v)
	}

	for k := range countryCapitalMap {
		fmt.Println("range2:", k, "-", countryCapitalMap[k])
	}

	capital, ok := countryCapitalMap["American"]

	if ok {
		fmt.Println("American -", capital)
	} else {
		fmt.Println("American not found")
	}
}

func TestMap2(t *testing.T) {

	countryCapitalMap := map[string]string{"France": "Paris", "Italy": "Rome", "Japan": "Tokyo", "India": "New delhi"}

	for k, v := range countryCapitalMap {
		fmt.Println("before:", k, "-", v)
	}

	delete(countryCapitalMap, "France")

	for k, v := range countryCapitalMap {
		fmt.Println("after:", k, "-", v)
	}

}
