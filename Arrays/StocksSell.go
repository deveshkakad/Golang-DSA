package main

import (
	"fmt"
	"math"
)

func main() {
	arr := []int{7, 1, 5, 3, 6, 4}

	fmt.Println(findMaxProfitFromStocks(arr))
}

func findMaxProfitFromStocks(arr []int) int {

	buyPrice := math.MaxInt
	maxProfit := 0

	for i := 0; i < len(arr); i++ {
		if buyPrice > arr[i] {
			buyPrice = arr[i]
		} else {
			profit := arr[i] - buyPrice
			if profit > maxProfit {
				maxProfit = profit
			}
		}
	}
	return maxProfit
}
