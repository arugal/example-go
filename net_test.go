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

func TestDialTCP(t *testing.T) {
	port := "8282"
	listen, err := listen(port)
	if err != nil {
		t.Error(err)
	}
	defer listen.Close()
	conn, err := net.Dial("tcp", "127.0.0.1:"+port)
	if err != nil {
		t.Error(err)
	}
	defer conn.Close()
	read := []byte{1}
	s, err := conn.Read(read)
	if err != nil {
		t.Error(err)
	}
	if s != 1 {
		t.Failed()
	}
	if read[0] != 0 {
		t.Failed()
	}
}

func listen(port string) (net.Listener, error) {
	listen, err := net.Listen("tcp", "127.0.0.1:"+
		port)
	if err != nil {
		return nil, err
	}

	go func() {
		for {
			conn, err := listen.Accept()
			if err != nil {
				continue
			}
			defer conn.Close()
			_, _ = conn.Write([]byte{0})
		}
	}()

	return listen, nil
}
