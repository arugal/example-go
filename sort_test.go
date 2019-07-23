package example_go

import (
	"math"
	"testing"
)

/**
see https://www.cnblogs.com/onepixel/p/7674659.html
*/

const arrLen = 100000

func newArrs() []int {
	arrs := make([]int, arrLen)
	for i := 0; i < arrLen; i++ {
		arrs[i] = arrLen - i
	}
	return arrs
}

func checkArrs(arrs []int, b *testing.B) {
	len := len(arrs) - 1
	for i := 0; i < len; i++ {
		if arrs[i] > arrs[i+1] {
			b.Fail()
		}
	}
}

// 冒泡排序
func BenchmarkBubbleSort(b *testing.B) {
	b.ResetTimer()
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

	checkArrs(arrs, b)
}

// 选择排序
func BenchmarkSelectionSort(b *testing.B) {
	b.ResetTimer()
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

	checkArrs(arrs, b)
}

// 插入排序
func BenchmarkInsertionSort(b *testing.B) {
	b.ResetTimer()
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
	checkArrs(arrs, b)
}

// 希尔排序
func BenchmarkShellSort(b *testing.B) {
	b.ResetTimer()
	arrs := newArrs()
	len := len(arrs)
	for gap := math.Floor(float64(len / 2)); gap > 0; gap = math.Floor(gap / 2) {
		for i := int(gap); i < len; i++ {
			j, iGap := i, int(gap)
			current := arrs[i]
			for j-iGap >= 0 && current < arrs[j-iGap] {
				arrs[j] = arrs[j-int(gap)]
				j = j - iGap
			}
			arrs[j] = current
		}
	}
	checkArrs(arrs, b)
}

// 归并排序
func BenchmarkMergeSort(b *testing.B) {
	b.ResetTimer()
	arrs := newArrs()
	len := len(arrs)
	if len > 1 {
		middle := int(math.Floor(float64(len / 2)))
		//left, right := make([]int, middle), make([]int, len-middle)
		//copy(left, arrs)
		//copy(right, arrs)
	}

	checkArrs(arrs, b)
}

func merge(left []int, right []int) {

}
