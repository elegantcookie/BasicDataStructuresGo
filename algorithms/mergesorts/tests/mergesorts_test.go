package tests

import (
	"BasicDataStructuresGo/algorithms/mergesorts/recursive/onebuffer_mergesort"
	"BasicDataStructuresGo/algorithms/mergesorts/recursive/simple_mergesort"
	"testing"
)

// Returns a reversed array
func getReversedArray(arr []int) []int {
	arrLen := len(arr)
	newArr := make([]int, arrLen)
	for i := 0; i < arrLen; i++ {
		newArr[i] = arr[arrLen-i-1]
	}
	return newArr
}

// Returns an ordered array from 0 to n
func getSortedArray(n int) []int {
	newArr := make([]int, n)
	for i := 0; i < n; i++ {
		newArr[i] = i
	}
	return newArr
}

// Checks if arrayA is equal to arrayB
func equal(arrayA []int, arrayB []int) bool {
	aLen := len(arrayA)
	bLen := len(arrayB)
	if aLen != bLen {
		return false
	}
	for i := 0; i < bLen; i++ {
		if arrayA[i] != arrayB[i] {
			return false
		}
	}
	return true
}

func TestRecursiveOnebufferMS(t *testing.T) {
	t.Run("RecursiveOnebufferMergeSort", func(t *testing.T) {

		t.Run("100 elems", func(t *testing.T) {
			sortedArr := getSortedArray(100)
			revArr := getReversedArray(sortedArr)
			onebuffer_mergesort.Sort(revArr)
			if !equal(sortedArr, revArr) {
				t.Errorf("WA")
			}
		})

		t.Run("10000 elems", func(t *testing.T) {
			sortedArr := getSortedArray(10000)
			revArr := getReversedArray(sortedArr)
			onebuffer_mergesort.Sort(revArr)
			if !equal(sortedArr, revArr) {
				t.Errorf("WA")
			}
		})

		t.Run("1000000 elems", func(t *testing.T) {
			sortedArr := getSortedArray(1000000)
			revArr := getReversedArray(sortedArr)
			onebuffer_mergesort.Sort(revArr)
			if !equal(sortedArr, revArr) {
				t.Errorf("WA")
			}
		})

		t.Run("10000000 elems", func(t *testing.T) {
			sortedArr := getSortedArray(10000000)
			revArr := getReversedArray(sortedArr)
			onebuffer_mergesort.Sort(revArr)
			if !equal(sortedArr, revArr) {
				t.Errorf("WA")
			}
		})

		t.Run("100000000 elems", func(t *testing.T) {
			sortedArr := getSortedArray(100000000)
			revArr := getReversedArray(sortedArr)
			onebuffer_mergesort.Sort(revArr)
			if !equal(sortedArr, revArr) {
				t.Errorf("WA")
			}
		})
	})
}

func TestRecursiveSimpleMS(t *testing.T) {
	t.Run("RecursiveSimpleMergeSort", func(t *testing.T) {

		t.Run("100 elems", func(t *testing.T) {
			sortedArr := getSortedArray(100)
			revArr := getReversedArray(sortedArr)
			simple_mergesort.Sort(revArr)
			if !equal(sortedArr, revArr) {
				t.Errorf("WA")
			}
		})

		t.Run("10000 elems", func(t *testing.T) {
			sortedArr := getSortedArray(10000)
			revArr := getReversedArray(sortedArr)
			simple_mergesort.Sort(revArr)
			if !equal(sortedArr, revArr) {
				t.Errorf("WA")
			}
		})

		t.Run("1000000 elems", func(t *testing.T) {
			sortedArr := getSortedArray(1000000)
			revArr := getReversedArray(sortedArr)
			simple_mergesort.Sort(revArr)
			if !equal(sortedArr, revArr) {
				t.Errorf("WA")
			}
		})

		t.Run("10000000 elems", func(t *testing.T) {
			sortedArr := getSortedArray(10000000)
			revArr := getReversedArray(sortedArr)
			simple_mergesort.Sort(revArr)
			if !equal(sortedArr, revArr) {
				t.Errorf("WA")
			}
		})

		t.Run("100000000 elems", func(t *testing.T) {
			sortedArr := getSortedArray(100000000)
			revArr := getReversedArray(sortedArr)
			simple_mergesort.Sort(revArr)
			if !equal(sortedArr, revArr) {
				t.Errorf("WA")
			}
		})
	})
}
