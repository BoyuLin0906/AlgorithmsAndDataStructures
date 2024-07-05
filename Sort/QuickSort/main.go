package main

import "fmt"

func partition(arr []int, start_idx int, end_idx int) int {
	pivot := arr[end_idx]

	j := start_idx - 1
	for i := start_idx; i < end_idx; i++ {
		if arr[i] < pivot {
			j += 1
			swap(&arr[i], &arr[j])
		}
	}
	j += 1
	swap(&arr[j], &arr[end_idx])
	return j
}

func swap(a *int, b *int) {
	*a, *b = *b, *a
}

func quickSort(arr []int, start_idx int, end_idx int) {
	if start_idx < end_idx {
		pivot_idx := partition(arr, start_idx, end_idx)
		quickSort(arr, start_idx, pivot_idx-1)
		quickSort(arr, pivot_idx+1, end_idx)
	}
}

func main() {
	arr := []int{9, 4, 1, 6, 7, 3, 8, 2, 10, 13, 50, 5}
	fmt.Println(arr)
	quickSort(arr, 0, len(arr)-1)
	fmt.Println(arr)
}
