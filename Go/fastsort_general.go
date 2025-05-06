package main

import "fmt"

var A [6]int = [6]int{49, 11, 35, 50, 34, 1}
var s []int = A[0:6] //slice切片，指向数组的指针？会传引用？

func main() {
	fastsort(s)
	fmt.Println(s)
}

func fastsort(A []int) {
	if len(A) < 2 {
		return
	}
	lowerA, upperA := partition(A)
	fastsort(lowerA)
	fastsort(upperA)
}

func partition(A []int) ([]int, []int) {
	var lower int = 0

	for i := 0; i < len(A); i++ {
		if A[i] < A[len(A)-1] {
			A[i], A[lower] = A[lower], A[i]
			lower++
		}
	}

	A[len(A)-1], A[lower] = A[lower], A[len(A)-1]

	return A[0:lower], A[lower+1 : len(A)]
}

/*
49 11 35 50 34 1
1 11 35 50 34 49 :11 35 50 34 49
11 35 34 49 50
11 35 34



*/
