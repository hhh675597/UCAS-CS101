package main

import "fmt"

func main() {
	//fmt.Printf("%c %c\n", 97, 65)
	var X byte = 62
	fmt.Println("X=", X, "\t-X=", -X, "\tX&1=", X&1, "\tX|1=", X|1, "\tX^1=", X^1)
}
