package main

import "fmt"

func sort(arr []int) []int {
	var i int
	var j int
	var size = len(arr)
	for i = 0; i < size-1; i++ {
		var min = i
		for j = i + 1; j < size; j++ {
			if arr[j] < arr[min] {
				min = j
			}
		}
		arr[i], arr[min] = arr[min], arr[i]
	}
	return arr
}

func main() {
	var a []int
	a = []int{5, 3, 4, 1, 2}
	fmt.Println(sort(a))
}
