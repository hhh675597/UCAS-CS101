package main

import "fmt"

func main() {
	var name string = "Huang Hao Huan"

	sum := 0
	for i := 0; i < len(name); i++ {
		sum = sum + int(name[i])
	}
	fmt.Println("sum = ", sum)
	/*
	   使用char类型输出

	   	var sum_bytes [4]byte
	   	var j int
	   	for j = 3; sum != 0; j-- {
	   		sum_bytes[j] = byte(sum%10) + '0'
	   		sum = sum / 10
	   	}

	   	for i := 0; i < 4; i++ {
	   		fmt.Printf("%c", sum_bytes[i])
	   	}
	   	fmt.Printf("\n")
	*/
}
