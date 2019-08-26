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

func TestDial(t *testing.T) {
	port := "8282"
	fetch := 9
	listen, err := listen(port, func(conn net.Conn) {
		defer conn.Close()
		_, _ = conn.Write([]byte{byte(fetch)})
	})
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
		t.Errorf("read err %v", err)
	}
	if s != 1 {
		t.Errorf("read result len = %d", s)
	}
	if len(read) < 1 || read[0] != byte(fetch) {
		t.Errorf("read %d != %d", read[0], fetch)
	}
}

func listen(port string, handler func(conn net.Conn)) (net.Listener, error) {
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
			handler(conn)
		}
	}()

	return listen, nil
}
