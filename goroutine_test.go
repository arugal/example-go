package example_go

import (
	"fmt"
	"github.com/prometheus/common/log"
	"net/http"
	"sync"
	"testing"
)

type worker struct {
	in   chan int
	done chan bool
}

func createWorker() worker {
	w := worker{
		in:   make(chan int),
		done: make(chan bool),
	}

	go func(w worker) {
		for n := range w.in {
			fmt.Println("rceived:", n)
			w.done <- true
		}
	}(w)

	return w
}

func TestChanNormal(t *testing.T) {
	var workers [10]worker

	for i, _ := range workers {
		workers[i] = createWorker()
	}

	for i, w := range workers {
		w.in <- i
	}

	for _, w := range workers {
		<-w.done
	}
}

func TestSyncWaitGroup(t *testing.T) {
	wg := sync.WaitGroup{}
	urls := []string{
		"https://www.baidu.com",
		"https://www.qq.com:",
	}

	for _, url := range urls {
		wg.Add(1)
		go func(url string) {
			defer wg.Done()
			resp, err := http.Get(url)
			if err != nil {
				log.Errorf("http get err %v", err)
				return
			}
			var body []byte
			n, err := resp.Body.Read(body)
			if err != nil {
				log.Errorf("body read err %v", err)
				return
			}
			log.Info(string(body[0:n]))
		}(url)
	}

	wg.Wait()
}
