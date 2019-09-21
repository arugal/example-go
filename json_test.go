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
