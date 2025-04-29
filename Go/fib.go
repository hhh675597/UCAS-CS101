package main

import "fmt"

var mem [31]int

func main() {

	for i := 0; i < 31; i++ {
		mem[i] = -1
	}

	fmt.Println(fib(30))
}

func fib(n int) int {
	fmt.Println("F(", n, ")")
	if n <= 1 {
		mem[n] = n
		return n
	}

	if mem[n] != -1 {
		fmt.Println("Immediate Return: F(", n, ") = ", mem[n])
		return mem[n]
	}

	mem[n] = fib(n-1) + fib(n-2)
	return mem[n]
}
