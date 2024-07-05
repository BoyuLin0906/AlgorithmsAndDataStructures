package main

import "fmt"

func maxHeapify(arr []int, idx int, length int) {
	left := idx*2 + 1
	right := idx*2 + 2

	largest := idx
	if left <= length && arr[idx] < arr[left] {
		largest = left
	}

	if right <= length && arr[largest] < arr[right] {
		largest = right
	}

	if idx != largest {
		swap(&arr[idx], &arr[largest])
		maxHeapify(arr, largest, length)
	}
}

func buildMaxHeap(arr []int) {
	for i := len(arr)/2 - 1; i > -1; i-- {
		maxHeapify(arr, i, len(arr)-1)
	}
}

func heapSort(arr []int) {
	buildMaxHeap(arr)
	size := len(arr) - 1
	for i := len(arr) - 1; i >= 1; i-- {
		size -= 1
		swap(&arr[0], &arr[i])
		maxHeapify(arr, 0, size)
	}
}

func swap(a *int, b *int) {
	*a, *b = *b, *a
}

func main() {
	arr := []int{9, 4, 1, 6, 7, 3, 8, 2, 5, 10, 13, 50, 5}
	fmt.Println(arr)
	heapSort(arr)
	fmt.Println(arr)
}
