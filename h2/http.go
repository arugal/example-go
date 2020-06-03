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
	"context"
	"crypto/tls"
	"fmt"
	"log"
	"net"
	"net/http"

	"golang.org/x/net/http2"
	"golang.org/x/net/http2/h2c"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "You tell %s\n", r.Proto)
	})
	h2s := &http2.Server{}
	h1s := &http.Server{Addr: ":9100", Handler: h2c.NewHandler(mux, h2s)}
	go func() {
		log.Fatal(h1s.ListenAndServe())
	}()
	defer h1s.Shutdown(context.Background())
	client := http.Client{
		// Skip TLS dial
		Transport: &http2.Transport{
			AllowHTTP: true,
			DialTLS: func(network, addr string, cfg *tls.Config) (net.Conn, error) {
				return net.Dial(network, addr)
			},
		},
	}
	resp, err := client.Get("http://127.0.0.1:9100")
	if err != nil {
		log.Fatal(fmt.Errorf("get response error: %v", err))
	}
	fmt.Println(resp.StatusCode)
	fmt.Println(resp.Proto)
}
