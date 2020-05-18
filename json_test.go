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
	"encoding/json"
	"fmt"
	"testing"
)

type S struct {
	Age int
	Map map[string]string
}

func TestJson1(t *testing.T) {
	m := make(map[string]string)
	m["1"] = "1"
	b, err := json.Marshal(&S{
		Age: 0,
		Map: m,
	})

	if err != nil {
		t.Error(err)
	}

	fmt.Println(string(b))
}
