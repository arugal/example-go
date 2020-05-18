// Copyright 2020 arugal, zhangwei24@apache.com
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

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
