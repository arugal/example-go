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
	"fmt"
	"log"
	"net/http"
	"sync"
	"time"

	"github.com/gin-gonic/gin"

	agent "github.com/SkyAPM/go2sky"
	plugin "github.com/SkyAPM/go2sky/plugins/gin"
	h "github.com/SkyAPM/go2sky/plugins/http"
	"github.com/SkyAPM/go2sky/reporter"
)

func main() {
	// Use gRPC reporter for production
	re, err := reporter.NewGRPCReporter("192.168.2.124:11800")
	if err != nil {
		log.Fatalf("new reporter error %v \n", err)
	}
	defer re.Close()

	tracer, err := agent.NewTracer("gin-server", agent.WithReporter(re))
	if err != nil {
		log.Fatalf("create tracer error %v \n", err)
	}
	tracer.WaitUntilRegister()
	gin.SetMode(gin.ReleaseMode)
	r := gin.New()

	//Use go2sky middleware with tracing
	r.Use(plugin.Middleware(r, tracer))

	r.GET("/user/:name", func(c *gin.Context) {
		name := c.Param("name")
		c.String(200, "Hello %s", name)
	})

	go func() {
		if err := http.ListenAndServe(":8080", r); err != nil {
			panic(err)
		}
	}()
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		defer wg.Done()
		for i := 0; i < 1000; i++ {
			request(tracer)
			time.Sleep(time.Millisecond * 100)
		}
	}()
	wg.Wait()
	time.Sleep(time.Second * 10)
	// Output:
}

func request(tracer *agent.Tracer) {
	//NewClient returns an HTTP Client with tracer
	client, err := h.NewClient(tracer)
	if err != nil {
		log.Fatalf("create client error %v \n", err)
	}
	request, err := http.NewRequest("GET", fmt.Sprintf("%s/user/gin", "http://127.0.0.1:8080"), nil)
	if err != nil {
		log.Fatalf("unable to create http request: %+v\n", err)
	}
	res, err := client.Do(request)
	if err != nil {
		log.Fatalf("unable to do http request: %+v\n", err)
	}
	_ = res.Body.Close()
}
