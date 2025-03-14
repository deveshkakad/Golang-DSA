package main

import "fmt"

func main() {
	arr := []int{3, 6, 2, 1, 8, 7, 4, 5, 3}

	fmt.Println("Array before sorting in Descending order:", arr)

	bubbleSortDesc(arr)

	fmt.Println("Array after sorting in Descending order:", arr)
}

func bubbleSortDesc(arr []int) {
	for i := 0; i < len(arr); i++ {
		for j := 0; j < len(arr)-1-i; j++ {
			if arr[j] < arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
	}
}
