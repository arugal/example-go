package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"sync"
	"time"

	agent "github.com/SkyAPM/go2sky"
	plugin "github.com/SkyAPM/go2sky/plugins/gin"
	h "github.com/SkyAPM/go2sky/plugins/http"
	"github.com/SkyAPM/go2sky/reporter"
)

func main() {
	// Use gRPC reporter for production
	re, err := reporter.NewLogReporter()
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
	r.GET("/user/:name", plugin.Middleware(r, tracer), func(c *gin.Context) {
		name := c.Param("name")
		c.String(200, "Hello %s", name)
	})

	//Use go2sky middleware with tracing
	r.Use(plugin.Middleware(r, tracer))

	go func() {
		if err := http.ListenAndServe(":8080", r); err != nil {
			panic(err)
		}
	}()
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		defer wg.Done()
		request(tracer)
	}()
	wg.Wait()
	time.Sleep(time.Second)
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
