package main

import "fmt"

func main() {
	arr := []int{3, 6, 2, 1, 8, 7, 4, 5, 3}

	fmt.Println("Array before sorting in Descending order:", arr)

	selectionSortDesc(arr)

	fmt.Println("Array after sorting in Descending order:", arr)
}

func selectionSortDesc(arr []int) {
	for i := 0; i < len(arr); i++ {
		for j := i + 1; j < len(arr); j++ {
			if arr[i] < arr[j] {
				arr[i], arr[j] = arr[j], arr[i]
			}
		}
	}
}
