package main

import "fmt"

func main() {
	arr := []int{3, 6, 2, 1, 8, 7, 4, 5, 3, -2}
	mergeSort(arr, 0, len(arr)-1)
	fmt.Println(arr)
}

func mergeSort(arr []int, start, end int) {
	if start >= end {
		return
	}

	mid := start + (end-start)/2

	mergeSort(arr, 0, mid)
	mergeSort(arr, mid+1, end)

	merge(arr, start, mid, end)
}

func merge(arr []int, start, mid, end int) {
	temp := make([]int, end-start+1)
	i, j, k := start, mid+1, 0

	for i <= mid && j <= end {
		if arr[i] < arr[j] {
			temp[k] = arr[i]
			i++
		} else {
			temp[k] = arr[j]
			j++
		}
		k++
	}

	for i <= mid {
		temp[k] = arr[i]
		k++
		i++
	}
	for j <= end {
		temp[k] = arr[j]
		k++
		j++
	}

	k = start
	for i = 0; i < len(temp); i++ {
		arr[k] = temp[i]
		k++
	}
}
