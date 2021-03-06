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
	"io/ioutil"
	"net/http"
	"testing"
)

func TestGet(t *testing.T) {
	resp, err := http.Get("http://www.baidu.com")
	if err != nil {
		t.Error(err)
	}
	defer resp.Body.Close()
	_, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		t.Error(err)
	}
}
