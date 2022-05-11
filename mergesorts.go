package main

import "fmt"

func merge(arr []int, l, mid, r int) {

	// 1 2 3 4 5
	// l = 0, r = 4, mid = 2

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

func mergeSort(arr []int, l, r int) {
	fmt.Printf("l: %v, r: %v\n", l, r)
	mid := (r + l) / 2
	if l == r {
		return
	}

	mergeSort(arr, l, mid)
	mergeSort(arr, mid+1, r)
	merge(arr, l, mid, r)

}
