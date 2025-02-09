package main

import "fmt"

func main() {
	arr := []int{1, 2, 3, 4, 5, 6, 7, 8}

	fmt.Print("Enter the element to be search in array:")
	var target int
	fmt.Scanln(&target)

	index := binarySearchRec(arr, 0, len(arr)-1, target)

	if index == -1 {
		fmt.Println("Element:", target, "is not present in the array")
	} else {
		fmt.Println("Element:", target, "is present in the array at index:", index)
	}

}

func binarySearchRec(arr []int, start, end, target int) int {
	if start > end {
		return -1
	}

	mid := (start + end) / 2

	if arr[mid] == target {
		return mid
	}

	if arr[mid] < target {
		return binarySearchRec(arr, mid+1, end, target)
	}

	return binarySearchRec(arr, start, mid-1, target)
}
