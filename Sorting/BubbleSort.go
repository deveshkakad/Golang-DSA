package main

import "fmt"

func main() {
	arr := []int{4, 2, 3, 5, 1}

	fmt.Println("Array before sorting:", arr)
	bubbleSort(arr)
	fmt.Println("Array after sorting:", arr)
}

func bubbleSort(arr []int) {
	for i := 0; i < len(arr); i++ {
		for j := 0; j < len(arr)-1-i; j++ {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
	}
}
