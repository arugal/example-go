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

package main

import (
	"bufio"
	"net"
	"net/http"

	frpIo "github.com/fatedier/golib/io"
	"github.com/sirupsen/logrus"
)

var log = logrus.New()

func echoServer() {
	l, err := net.Listen("tcp", ":1234")
	if err != nil {
		panic(err)
	}
	defer l.Close()
	buf := make([]byte, 8)
	for {
		conn, err := l.Accept()
		if err != nil {
			log.Warnf("echo accept error: %v", err)
			continue
		}
		go func() {
			defer conn.Close()
			n, err := conn.Read(buf)
			if err != nil {
				log.Warnf("echo read conn error: %v", err)
				return
			}
			_, _ = conn.Write(buf[:n])
		}()
	}
}

func httpServer() {
	http.HandleFunc("/v1", func(writer http.ResponseWriter, request *http.Request) {
		_, _ = writer.Write([]byte("hello world!"))
		writer.WriteHeader(http.StatusOK)
	})
	err := http.ListenAndServe(":1235", nil)
	if err != nil {
		panic(err)
	}
}

func dispatchServer() {
	l, err := net.Listen("tcp", ":1233")
	if err != nil {
		panic(err)
	}
	defer l.Close()
	for {
		conn, err := l.Accept()
		if err != nil {
			log.Warnf("dispatch accept error: %v", err)
		}
		go func() {
			defer conn.Close()
			buffer := bufio.NewReader(conn)
			var err error
			var head []byte
			head, err = buffer.Peek(3)
			if err != nil {
				log.Warnf("dispatch peek conn error: %v", err)
				return
			}
			var targetConn net.Conn
			if "GET" == string(head) {
				// http
				targetConn, err = net.Dial("tcp", "127.0.0.1:1235")
				if err != nil {
					log.Warnf("dial http server error: %v", err)
					return
				}
			} else {
				// echo
				targetConn, err = net.Dial("tcp", "127.0.0.1:1234")
				if err != nil {
					log.Warnf("dial echo server error: %v", err)
					return
				}
			}
			defer targetConn.Close()
			frpIo.Join(targetConn, conn)
		}()
	}
}

func main() {
	go echoServer()
	go httpServer()
	dispatchServer()
}
