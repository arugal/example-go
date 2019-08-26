package example_go

import (
	"net"
	"testing"
)

func TestResolveIPAddr(t *testing.T) {
	_, err := net.ResolveIPAddr("ip", "www.baidu.com")

	if err != nil {
		t.Error(err)
	}
}
