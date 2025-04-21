package main

var mem [31]int

func main() {
	for i := 0; i < 31; i++ {
		mem[i] = -1
	}
	println(fib(30))
}

func fib(n int) int {
	if n <= 1 {
		return n
	}
	if mem[n] != -1 {
		return mem[n]
	}
	mem[n] = fib(n-1) + fib(n-2)
	return mem[n]
}
