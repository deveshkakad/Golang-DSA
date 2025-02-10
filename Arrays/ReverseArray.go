package main

import "fmt"

func main() {
	arr := []int{1, 2, 3, 4, 5}

	fmt.Print("Original Array:")
	fmt.Println(arr)

	reverseArray(arr)

	fmt.Print("Reversed Array:")
	fmt.Println(arr)
}

func reverseArray(arr []int) {
	start := 0
	end := len(arr) - 1

	for start < end {
		temp := arr[start]
		arr[start] = arr[end]
		arr[end] = temp

		start++
		end--
	}
}
