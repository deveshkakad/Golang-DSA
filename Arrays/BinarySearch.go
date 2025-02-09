package main

import "fmt"

func main() {
	arr := []int{1, 2, 3, 4, 5, 6, 7, 8}

	fmt.Print("Enter the element to be search in array:")
	var target int
	fmt.Scanln(&target)

	index := binarySearch(arr, target)

	if index == -1 {
		fmt.Println("Element:", target, "is not present in the array")
	} else {
		fmt.Println("Element:", target, "is present in the array at index:", index)
	}
}
func binarySearch(arr []int, target int) int {
	start := 0
	end := len(arr) - 1

	for start <= end {
		mid := (start + end) / 2

		if arr[mid] == target {
			return mid
		} else if arr[mid] < target {
			start = mid + 1
		} else {
			end = mid - 1
		}
	}
	return -1
}
