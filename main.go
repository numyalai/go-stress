package main

import "fmt"

func BubbleSort(arr []int) []int {
	var n = len(arr)
	for i := 0; i < n; i++ {
		for j := 0; j < n-1; j++ {
			if arr[j] > arr[j+1] {
				var temp = arr[j]
				arr[j] = arr[j+1]
				arr[j+1] = temp
			}
		}
	}
	return arr
}

func SelectionSort(arr []int) []int {
	var n = len(arr)
	for i := 0; i < n; i++ {
		var min = i
		for j := i + 1; j < n; j++ {
			if arr[j] < arr[min] {
				min = j
			}
		}
		if min != i {
			var temp = arr[i]
			arr[i] = arr[min]
			arr[min] = temp
		}
	}
	return arr
}

func InsertionSort(arr []int) []int {
	var n = len(arr)
	for i := 1; i < n; i++ {
		var key = arr[i]
		var j = i - 1
		for j >= 0 && arr[j] > key {
			arr[j+1] = arr[j]
			j = j - 1
		}
		arr[j+1] = key
	}
	return arr
}

func main() {
	var arr = []int{3, 44, 38, 5, 47, 15, 36, 26, 27, 2, 46, 4, 19, 50, 48}
	SelectionSort(arr)

	fmt.Println(arr)
}
