package main

import "fmt"

func main() {
	arr := []int{1, 2, 32, 21, 45, 31}

	var target int
	fmt.Print("Enter the target element to be search:")
	fmt.Scanln(&target)
	index := linearSearch(arr, target)

	if index == -1 {
		fmt.Println("Element ", target, "is not present in the array")
	} else {
		fmt.Println("Element", target, "is present in the array at index:", index)
	}
}

func linearSearch(arr []int, target int) int {

	for i, v := range arr {
		if v == target {
			return i
		}
	}
	return -1
}
