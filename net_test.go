package example_go

import (
	"net"
	"testing"
)

func TestResolveIPAddr(t *testing.T) {
	_, err := net.ResolveIPAddr("ip", "www.sunnus3.online")

	if err != nil {
		t.Error(err)
	}
}
