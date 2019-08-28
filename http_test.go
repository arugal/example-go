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
