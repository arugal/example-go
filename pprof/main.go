package main

import (
	"fmt"
	"net/http"
	_ "net/http/pprof"
	"time"
)

func main() {
	go func() {
		for {
			fmt.Println("go")
			time.Sleep(time.Second)
		}
	}()
	_ = http.ListenAndServe("localhost:6060", nil)
}
