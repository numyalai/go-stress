package main

import (
	"fmt"
	"math/rand"
	"net/http"
)

func BubbleSort(arr []uint64) []uint64 {
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

func SelectionSort(arr []uint64) []uint64 {
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

func InsertionSort(arr []uint64) []uint64 {
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

func generateRandomArray() []uint64 {
	arr := make([]uint64, 10000)

	for i := range arr {
		arr[i] = uint64(rand.Uint64())
	}

	fmt.Println("First 10 elements of the array:", arr[:10])
	return arr
}

func main() {

	server := http.NewServeMux()

	server.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {

		go func() {
			arr := generateRandomArray()
			BubbleSort(arr)
		}()

		w.Header().Set("Content-Type", "text/plain")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Bubble sort done"))
	})

	http.ListenAndServe(":5010", server)

}
