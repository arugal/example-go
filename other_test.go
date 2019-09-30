package example_go

import (
	"fmt"
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
