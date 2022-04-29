package main

import "fmt"

func main() {
	arr := []int{5, 2, 7, 9, 11, 1}
	fastSort(arr, 0, len(arr) - 1)
	fmt.Println(arr)
}

func fastSort(arr []int, l, r int) {
	if l >= r {
		return
	}
	m := loop(arr, l, r)
	fastSort(arr, l, m-1)
	fastSort(arr, m+1, r)
}

func loop(arr []int, l, r int) int {
	i, pivot := l, arr[r]
	for j := l;j < r;j++ {
		if arr[j] <= pivot {
			arr[i], arr[j] = arr[j], arr[i]
			i++
		}
	}
	arr[i], arr[r] = arr[r], arr[i]
	return i
}