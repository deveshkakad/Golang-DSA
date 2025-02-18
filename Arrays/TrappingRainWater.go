package main

import (
	"fmt"
)

func main() {
	arr := []int{4, 2, 0, 6, 3, 2, 5}

	ans := getTrappingWater(arr)

	fmt.Println("WaterLevel trapped:", ans)
}

func getTrappingWater(arr []int) int {
	waterLevel := 0
	n := len(arr)
	left := make([]int, n)
	right := make([]int, n)

	left[0] = arr[0]
	for i := 1; i < n; i++ {
		left[i] = getMax(left[i-1], arr[i])
	}

	right[n-1] = arr[n-1]
	for i := n - 2; i >= 0; i-- {
		right[i] = getMax(right[i+1], arr[i])
	}

	for i := 0; i < n; i++ {
		waterTrapped := getMin(left[i], right[i]) - arr[i]
		waterLevel += waterTrapped
	}

	return waterLevel
}

func getMin(n1 int, n2 int) int {
	if n1 < n2 {
		return n1
	} else {
		return n2
	}
}
func getMax(n1 int, n2 int) int {
	if n1 > n2 {
		return n1
	} else {
		return n2
	}
}
