package main

import "fmt"

func main() {
	arr := []int{1, 2, 3, 4, 5}

	fmt.Print("Original Array:")
	fmt.Println(arr)

	reverseArrayRec(arr, 0, len(arr)-1)

	fmt.Print("Reversed Array:")
	fmt.Println(arr)
}

func reverseArrayRec(arr []int, start, end int) {
	if start >= end {
		return
	}

	temp := arr[start]
	arr[start] = arr[end]
	arr[end] = temp

	reverseArrayRec(arr, start+1, end-1)
}
