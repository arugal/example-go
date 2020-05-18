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
	"sync"
	"testing"
)

func TestLabelWith(t *testing.T) {
	continueInx := 0
	gotoInx := 0
	breakInx := 0
loop:
	for {
		continueInx++
		gotoInx++
		breakInx++
		if continueInx > 3 {
			continueInx = 0
			continue loop
		}
		if gotoInx > 5 {
			gotoInx = 0
			goto loop
		}
		if breakInx > 10 {
			break loop
		}
		fmt.Printf("continueInx:%d, gotoInx:%d, breakInx:%d \n", continueInx, gotoInx, breakInx)
	}
}

type Point struct {
	Inx1 int
	Inx2 *int
}

func (p Point) func1() {
	p.Inx1 = p.Inx1 + 1
	inx := *p.Inx2 + 1
	p.Inx2 = &inx
	fmt.Printf("In func1 Inx1 %d, Inx2 %d \n", p.Inx1, *p.Inx2)
}

func (p *Point) func2() {
	p.Inx1 = p.Inx1 + 1
	inx := *p.Inx2 + 1
	p.Inx2 = &inx
	fmt.Printf("In func1 Inx1 %d, Inx2 %d \n", p.Inx1, *p.Inx2)
}

func TestPoint(t *testing.T) {
	inx2 := 1
	point1 := Point{
		Inx1: 1,
		Inx2: &inx2,
	}
	point1.func1()
	fmt.Println(point1)

	inx22 := 1
	point2 := Point{
		Inx1: 1,
		Inx2: &inx22,
	}
	point2.func2()
	fmt.Println(point2)
}

func TestSyncMap(t *testing.T) {
	m := sync.Map{}

	done := make(chan bool)

	go func() {
		m.Store("abc", "def")
		done <- true
	}()

	<-done
	m.Range(func(key, value interface{}) bool {
		fmt.Printf("key: %v, value %v \n", key, value)
		return true
	})

	m.Delete("abc")

	m.Range(func(key, value interface{}) bool {
		fmt.Printf("key: %v, value %v", key, value)
		return true
	})

}
