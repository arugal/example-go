package example_go

import (
	"bufio"
	"crypto/tls"
	"fmt"
	"net"
	"testing"
	"time"

	"github.com/hashicorp/yamux"
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

func TestTls(t *testing.T) {
	port := "8283"
	listen, err := listenTls(port, "certs/server.pem", "certs/server.key", func(conn net.Conn) {
		defer conn.Close()

		r := bufio.NewReader(conn)

		for {
			msg, err := r.ReadString('\n')
			if err != nil {
				return
			}
			fmt.Println("listen: " + msg)

			_, err = conn.Write([]byte("world\n"))
			if err != nil {
				return
			}
		}
	})

	if err != nil {
		t.Error(err)
	}
	defer listen.Close()

	conf := &tls.Config{
		InsecureSkipVerify: true,
	}
	conn, err := tls.Dial("tcp", "127.0.0.1:"+port, conf)
	if err != nil {
		t.Error(err)
	}
	defer conn.Close()

	n, err := conn.Write([]byte("hello\n"))
	if err != nil {
		t.Error(err)
	}
	buf := make([]byte, 100)
	n, err = conn.Read(buf)
	if err != nil {
		t.Error(err)
	}
	fmt.Println(string(buf[:n]))
}

func listenTls(port string, certFile string, keyFile string, handler func(conn net.Conn)) (net.Listener, error) {

	cert, err := tls.LoadX509KeyPair(certFile, keyFile)

	if err != nil {
		return nil, err
	}

	config := &tls.Config{Certificates: []tls.Certificate{cert}}
	ln, err := tls.Listen("tcp", ":"+port, config)

	if err != nil {
		return nil, err
	}

	go func() {
		for {
			conn, err := ln.Accept()
			if err != nil {
				continue
			}
			go handler(conn)
		}
	}()
	return ln, nil
}

func TestYamux(t *testing.T) {
	port := "8283"

	listen, err := listenYamux(port, func(conn net.Conn) {
		defer conn.Close()
		fmt.Println("yamux")
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

	session, err := yamux.Client(conn, nil)
	if err != nil {
		t.Error(err)
	}
	defer session.Close()

	stream, err := session.Open()
	if err != nil {
		t.Error(err)
	}
	defer stream.Close()

	stream1, _ := session.Open()
	if err != nil {
		t.Error(err)
	}
	defer stream1.Close()
	time.Sleep(1 * time.Second)
}

func listenYamux(port string, handler func(conn net.Conn)) (net.Listener, error) {
	tcpaddr, _ := net.ResolveTCPAddr("tcp4", "127.0.0.1:"+port)
	listen, err := net.ListenTCP("tcp", tcpaddr)
	if err != nil {
		return nil, err
	}

	go func() {
		for {
			conn, err := listen.Accept()
			if err != nil {
				continue
			}
			session, err := yamux.Server(conn, nil)
			if err != nil {
				continue
			}
			go func() {
				defer session.Close()
				defer conn.Close()

				for {
					stream, err := session.Accept()
					if err != nil {
						return
					}
					go handler(stream)
				}
			}()
		}
	}()

	return listen, nil
}
