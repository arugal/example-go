package example_go

import "testing"
import "github.com/stretchr/testify/assert"

/**
see https://www.cnblogs.com/onepixel/p/7674659.html
*/

func newArrs() []int {
	arrs := make([]int, 10000)
	for i := 0; i < 10000; i++ {
		arrs[i] = 10000 - i
	}
	return arrs
}

func checkArrs(arrs []int) bool {
	len := len(arrs) - 1
	for i := 0; i < len; i++ {
		if arrs[i] > arrs[i+1] {
			return false
		}
	}
	return true
}

func TestBubbleSort(t *testing.T) {
	assert := assert.New(t)

	arrs := newArrs()
	len := len(arrs) - 1
	for i := 0; i < len; i++ {
		for j := 0; j < len; j++ {
			if arrs[j] > arrs[j+1] {
				temp := arrs[j+1]
				arrs[j+1] = arrs[j]
				arrs[j] = temp
			}
		}
	}

	assert.Equal(true, checkArrs(arrs))
}

func TestSelectionSort(t *testing.T) {
	assert := assert.New(t)

	arrs := newArrs()
	len := len(arrs)
	mixIndex, temp := 0, 0

	for i := 0; i < len-1; i++ {
		mixIndex = i
		for j := i + 1; j < len; j++ {
			if arrs[j] < arrs[mixIndex] {
				mixIndex = j
			}
		}
		temp = arrs[i]
		arrs[i] = arrs[mixIndex]
		arrs[mixIndex] = temp
	}

	assert.Equal(true, checkArrs(arrs))
}

func TestInsertionSort(t *testing.T) {
	assert := assert.New(t)

	arrs := newArrs()
	len := len(arrs)
	preIndex, current := 0, 0

	for i := 1; i < len; i++ {
		preIndex = i - 1
		current = arrs[i]

		for preIndex >= 0 && arrs[preIndex] > current {
			arrs[preIndex+1] = arrs[preIndex]
			preIndex--
		}
		arrs[preIndex+1] = current
	}

	assert.Equal(true, checkArrs(arrs))
}

func TestShellSort(t *testing.T) {
	assert := assert.New(t)

	arrs := newArrs()

	assert.Equal(true, checkArrs(arrs))
}
