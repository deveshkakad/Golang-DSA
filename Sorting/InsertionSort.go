package main

import "fmt"

func main() {
	arr := []int{3, 2, 1, 5, 4}

	fmt.Println("Array before sorting:", arr)
	insertionSort(arr)
	fmt.Println("Array after sorting:", arr)
}

func insertionSort(arr []int) {

	for i := 1; i < len(arr); i++ {
		curr := arr[i]
		j := i - 1
		for ; j >= 0 && arr[j] > curr; j-- {
			arr[j+1] = arr[j]
		}
		arr[j+1] = curr
	}
}
