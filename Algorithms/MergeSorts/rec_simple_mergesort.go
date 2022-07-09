package MergeSorts

// Merges left and right halves of the array with memory allocation for an array of the size of each half.
func merge(arr []int, l, mid, r int) {
	subArrayOne := mid - l + 1
	subArrayTwo := r - mid

	leftArr := make([]int, subArrayOne)
	rightArr := make([]int, subArrayTwo)

	for i := 0; i < subArrayOne; i++ {
		leftArr[i] = arr[l+i]
	}
	for i := 0; i < subArrayTwo; i++ {
		rightArr[i] = arr[mid+i+1]
	}
	i, j, k := 0, 0, l
	for i < subArrayOne && j < subArrayTwo {
		if leftArr[i] <= rightArr[j] {
			arr[k] = leftArr[i]
			i++
			k++
		} else {
			arr[k] = rightArr[j]
			j++
			k++
		}
	}
	for i < subArrayOne {
		arr[k] = leftArr[i]
		i++
		k++
	}

	for j < subArrayTwo {
		arr[k] = rightArr[j]
		j++
		k++
	}

}

// Simple mergesort. To sort an array, run it as mergeSort(array_ref, 0, n-1), where n is the size of the array
func mergeSort(array []int, l, r int) {

	mid := (r + l) / 2
	if l == r {
		return
	}

	mergeSort(array, l, mid)
	mergeSort(array, mid+1, r)
	merge(array, l, mid, r)

}
