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

package example_go

import (
	"fmt"
	"net/http"
	"sync"
	"testing"

	"github.com/prometheus/common/log"
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
