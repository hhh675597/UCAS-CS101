package main // compare.go
import "fmt"

func main() {
	var X byte = 62
	fmt.Println("X=", X, "\t-X=", -X, "\tX&1=", X&1, "\tX|1=", X|1, "\tX^1=", X^1)
}
