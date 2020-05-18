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

package template

import (
	"os"
	"testing"
	"text/template"
)

type Friend struct {
	Name string
}

type Person struct {
	UserName string
	Emails   []string
	Friends  []*Friend
}

func Test1(t *testing.T) {
	f1 := Friend{Name: "Zhang San"}
	f2 := Friend{Name: "Li Si"}
	temp := template.New("test1")
	temp = template.Must(temp.Parse(
		`hello {{.UserName}}!
{{ range .Emails }}
an email {{ . }}
{{- end }}
{{ with .Friends }}
{{- range . }}
my friend name is {{.Name}}
{{- end }}
{{ end }}`))
	p := Person{UserName: "Wang Wu",
		Emails:  []string{"wangwu@qq.com", "wangwu@gmail.com"},
		Friends: []*Friend{&f1, &f2}}
	temp.Execute(os.Stdout, p)
}
