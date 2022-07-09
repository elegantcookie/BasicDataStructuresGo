package simple_mergesort

import (
	"BasicDataStructuresGo/misc/array_tests"
	"testing"
)

func TestRecursiveSimpleMS(t *testing.T) {
	t.Run("RecursiveSimpleMergeSort", func(t *testing.T) {

		t.Run("100 elems", func(t *testing.T) {
			sortedArr := array_tests.GetSortedArray(100)
			revArr := array_tests.GetReversedArray(sortedArr)
			Sort(revArr)
			if !array_tests.Equal(sortedArr, revArr) {
				t.Errorf("WA")
			}
		})

		t.Run("10000 elems", func(t *testing.T) {
			sortedArr := array_tests.GetSortedArray(10000)
			revArr := array_tests.GetReversedArray(sortedArr)
			Sort(revArr)
			if !array_tests.Equal(sortedArr, revArr) {
				t.Errorf("WA")
			}
		})

		t.Run("1000000 elems", func(t *testing.T) {
			sortedArr := array_tests.GetSortedArray(1000000)
			revArr := array_tests.GetReversedArray(sortedArr)
			Sort(revArr)
			if !array_tests.Equal(sortedArr, revArr) {
				t.Errorf("WA")
			}
		})

		t.Run("10000000 elems", func(t *testing.T) {
			sortedArr := array_tests.GetSortedArray(10000000)
			revArr := array_tests.GetReversedArray(sortedArr)
			Sort(revArr)
			if !array_tests.Equal(sortedArr, revArr) {
				t.Errorf("WA")
			}
		})

		t.Run("100000000 elems", func(t *testing.T) {
			sortedArr := array_tests.GetSortedArray(100000000)
			revArr := array_tests.GetReversedArray(sortedArr)
			Sort(revArr)
			if !array_tests.Equal(sortedArr, revArr) {
				t.Errorf("WA")
			}
		})
	})
}
