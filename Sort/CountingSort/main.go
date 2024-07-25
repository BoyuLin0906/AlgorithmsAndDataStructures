package main

import "fmt"

func countingSort(arr []int) []int {
	max_val := arr[0]
	min_val := arr[0]

	for i := 1; i < len(arr); i++ {
		if arr[i] > max_val {
			max_val = arr[i]
		}
		if arr[i] < min_val {
			min_val = arr[i]
		}
	} 

	// for negative value
	offset := 0
	if min_val < 0 {
		offset = -min_val
	}

	count := make([]int, max_val+offset+1)
	for i := 0; i < len(arr); i++ {
		count[arr[i]+offset] += 1
	}

	for i := 1; i < len(count); i++ {
		count[i] += count[i-1]
	} 

	output_arr := make([]int, len(arr))
	for i := 0; i < len(arr); i++ {
		count[arr[i]+offset] -= 1
		output_arr[count[arr[i]+offset]] = arr[i]
	}

	return output_arr
}

func main() {
	arr := []int{9, 3, 8, 10, 56, 27, 98, 1, 47, 12, 2, 7}
	fmt.Println(arr)
	output_arr := countingSort(arr)
	fmt.Println(output_arr)
}
