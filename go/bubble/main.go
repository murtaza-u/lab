package main

import "fmt"

func sort(arr []int) {
	n := len(arr)
	for i := 0; i < n-1; i++ {
		for j := 0; j < n-i-1; j++ {
			if arr[j] > arr[j+1] {
				temp := arr[j]
				arr[j] = arr[j+1]
				arr[j+1] = temp
			}
		}
	}
}

func main() {
	arr := []int{-2, 1, 190, 0, 4, -53, 21}
	fmt.Printf("unsorted: %v\n", arr)
	sort(arr)
	fmt.Printf("sorted: %v\n", arr)
}
