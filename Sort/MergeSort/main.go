package main

import (
	"fmt"
	"math"
)

func mergeSort(arr []int, start int, end int) {
	if start < end {
		middle := start + (end-start)/2
		mergeSort(arr, start, middle)
		mergeSort(arr, middle+1, end)
		merge(arr, start, middle+1, end)
	}
}

func merge(arr []int, start int, middle int, end int) {
	sub_left := make([]int, len(arr[start:middle]), len(arr[start:middle])+1)
	copy(sub_left, arr[start:middle])
	sub_left = append(sub_left, math.MaxInt)
	left_idx := 0

	sub_right := make([]int, len(arr[middle:end+1]), len(arr[middle:end+1])+1)
	copy(sub_right, arr[middle:end+1])
	sub_right = append(sub_right, math.MaxInt)
	right_idx := 0

	for i := start; i <= end; i++ {
		if sub_left[left_idx] < sub_right[right_idx] {
			arr[i] = sub_left[left_idx]
			left_idx += 1
		} else {
			arr[i] = sub_right[right_idx]
			right_idx += 1
		}
	}
}

func main() {
	arr := []int{9, 3, 8, 10, 56, 27, 98, 1, 47, 12, 2, 7}
	fmt.Println(arr)
	mergeSort(arr, 0, len(arr)-1)
	fmt.Println(arr)
}
