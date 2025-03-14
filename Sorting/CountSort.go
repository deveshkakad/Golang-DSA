package main

import (
	"fmt"
	"math"
)

func main() {
	arr := []int{1, 4, 1, 3, 2, 4, 3, 7}
	fmt.Println("Array before sorting:", arr)

	countArr(arr)

	fmt.Println("Array after sorting:", arr)
}

func countArr(arr []int) {
	maxi := math.MinInt
	j := 0
	for i := 0; i < len(arr); i++ {
		if arr[i] > maxi {
			maxi = arr[i]
		}
	}

	count := make([]int, maxi+1)
	for i := 0; i < len(arr); i++ {
		count[arr[i]]++
	}

	for i := 0; i < len(count); i++ {
		for count[i] > 0 {
			arr[j] = i
			count[i]--
			j++
		}
	}
}
