package array_tests

// Returns a reversed array
func GetReversedArray(arr []int) []int {
	arrLen := len(arr)
	newArr := make([]int, arrLen)
	for i := 0; i < arrLen; i++ {
		newArr[i] = arr[arrLen-i-1]
	}
	return newArr
}

// Returns an ordered array from 0 to n
func GetSortedArray(n int) []int {
	newArr := make([]int, n)
	for i := 0; i < n; i++ {
		newArr[i] = i
	}
	return newArr
}

// Checks if arrayA is equal to arrayB
func Equal(arrayA []int, arrayB []int) bool {
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
