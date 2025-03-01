package main

import "fmt"

func main() {
	arr := [3]int{3, 1, 2}
	res := make([]int, 0)
	subSequence(arr, res, 0)

}

func subSequence(arr [3]int, res []int, index int) {
	if index == len(arr) {
		fmt.Println(res)
		//printArr(res)
		return
	}

	res = append(res, arr[index])
	subSequence(arr, res, index+1)

	res = res[:len(res)-1]
	subSequence(arr, res, index+1)
}

func printArr(arr []int) {
	for _, v := range arr {
		fmt.Print(v, " ")
	}
	fmt.Println()
}
