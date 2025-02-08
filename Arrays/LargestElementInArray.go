package main

import (
	"fmt"
	"math"
)

func main() {
	arr := []int{1, 32, 21, 345, 234, 563, 23}

	index := largestElement(arr)

	fmt.Println("Largest element in the given array is:", arr[index], "present at index:", index)
}

func largestElement(arr []int) int {
	maxi := -1
	maxElement := math.MinInt
	for i, v := range arr {
		if v > maxElement {
			maxi = i
			maxElement = v
		}
	}

	return maxi
}
