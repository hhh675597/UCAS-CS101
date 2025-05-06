package main

import (
	"fmt"
	"math/rand"
)

func main() {
	var data [1024 * 1024]int
	for i := 0; i < 1024*1024; i++ {
		data[i] = i
	}
	var s []int = data[0 : 1024*1024]
	//var s []int = []int{4, 2, 1, 3, 5}
	quicksort(s)
	fmt.Println(s)
}

func quicksort(A []int) {
	if len(A) < 2 {
		return
	}
	lowerA, upperA := partition(A)
	quicksort(lowerA)
	quicksort(upperA)
}

func partition(A []int) ([]int, []int) {
	// 随机选取pivot
	pivotIndex := rand.Intn(len(A))
	A[pivotIndex], A[len(A)-1] = A[len(A)-1], A[pivotIndex]

	var lower int = 0
	for i := 0; i < len(A)-1; i++ {
		if A[i] < A[len(A)-1] {
			A[i], A[lower] = A[lower], A[i]
			lower++
		}
	}
	A[lower], A[len(A)-1] = A[len(A)-1], A[lower]
	return A[0:lower], A[lower+1 : len(A)]
}
