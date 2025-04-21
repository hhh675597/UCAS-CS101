package main

import "fmt"

func parity(x int) int {
	var parityValue int = 0

	for i := 0; i < 64; i++ {
		if x&1 == 1 {
			parityValue++
		}
		x >>= 1
	}
	return parityValue
}

func main() {
	var name string = "Huang Hao Huan"

	var sum int = 0
	for i := 0; i < len(name); i++ {
		sum = sum + int(name[i])
	}
	fmt.Println(parity(sum))
}
