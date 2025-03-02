package main

import "fmt"

func main() {
	arr := []int{3, 2, 1, 5, 4}

	fmt.Println("Array before sorting:", arr)
	selectionSort(arr)
	fmt.Println("Array after sorting:", arr)
}

func selectionSort(arr []int) {
	for i := 0; i < len(arr); i++ {
		for j := i + 1; j < len(arr); j++ {
			if arr[i] > arr[j] {
				arr[i], arr[j] = arr[j], arr[i]
			}
		}
	}
}
