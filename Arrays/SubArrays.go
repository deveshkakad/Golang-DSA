package main

import "fmt"

func main() {
	arr := []int{2, 4, 6, 8, 10}

	printSubArrays(arr)
}

func printSubArrays(arr []int) {
	n := len(arr)

	for i := 0; i < n; i++ {
		fmt.Println(arr[i])
		for j := i + 1; j < n; j++ {
			for start := i; start <= j; start++ {
				fmt.Print(arr[start], " ")
			}
			fmt.Println()
		}
	}
}
