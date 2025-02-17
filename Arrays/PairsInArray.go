package main

import "fmt"

func main() {
	arr := []int{2, 4, 6, 8, 10}

	printPairsInArray(arr)
}

func printPairsInArray(arr []int) {
	for i := 0; i < len(arr); i++ {
		for j := i + 1; j < len(arr); j++ {
			fmt.Print("(", arr[i], ",", arr[j], ")")
		}
		fmt.Println()
	}
}
