package main

import "fmt"

func insertionSort(arr []int) {
	for i := 1; i < len(arr); i++ {
		cur_val := arr[i]

		j := i - 1
		for j >= 0 && arr[j] > cur_val {
			arr[j+1] = arr[j]
			j -= 1
		}

		arr[j+1] = cur_val
	}
}

func main() {
	arr := []int{9, 3, 8, 10, 56, 27, 98, 1, 47, 12, 2, 7}
	fmt.Println(arr)
	insertionSort(arr)
	fmt.Println(arr)
}
