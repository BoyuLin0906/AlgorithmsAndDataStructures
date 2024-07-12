package main

import "fmt"

func bubbleSort(arr []int) {
	for i := 0; i < len(arr)-1; i++ {
		flag := false
		for j := 0; j < len(arr)-1-i; j++ {
			if arr[j] > arr[j+1] {
				swap(&arr[j], &arr[j+1])
				flag = true
			}
		}
		if !flag {
			fmt.Println(i)
			break
		}
	}
}

func swap(a *int, b *int) {
	*a, *b = *b, *a
}

func main() {
	arr := []int{9, 4, 1, 6, 7, 3, 8, 2, 5, 10, 13, 50, 5, 9, 82, 24, 39}
	fmt.Println(arr)
	bubbleSort(arr)
	fmt.Println(arr)
}
