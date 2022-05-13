package main

// Merges left and right halves of the array using buffer
func obMerge(arr []int, buffer []int, l, mid, r int) {
	subArrayOne := mid + 1
	subArrayTwo := r + 1

	for i := l; i < r+1; i++ {
		buffer[i] = arr[i]
	}

	i, j, k := l, subArrayOne, l

	for i < subArrayOne && j < subArrayTwo {
		if buffer[i] <= buffer[j] {
			arr[k] = buffer[i]
			i++
			k++
		} else {
			arr[k] = buffer[j]
			j++
			k++
		}
	}

	for i < subArrayOne {
		arr[k] = buffer[i]
		i++
		k++
	}

	for j < subArrayTwo {
		arr[k] = buffer[j]
		j++
		k++
	}
}

// Uses only one buffer instead of allocating memory each time the obMerge func is called.
// To sort an array, run it as mergeSort(array_ref, buffer_ref, 0, n-1), where n is the size of the array.
// Size of the arr and buffer must be the same.
func obMergeSort(array []int, buffer []int, l, r int) {

	mid := l + (r-l)/2
	if l == r {
		return
	}

	obMergeSort(array, buffer, l, mid)
	obMergeSort(array, buffer, mid+1, r)
	obMerge(array, buffer, l, mid, r)

}
