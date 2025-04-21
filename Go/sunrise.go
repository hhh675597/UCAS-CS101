package main

import (
	"fmt"
	"math"
)

func main() {
	var D float64
	B := 6*3600 + 2*60 //  2024.9.22 北京日出时间 6:02
	length := 40000 * math.Cos(40*3.14159/180)
	S := length / (24 * 3600)

	D = length * (116 - 113) / 360
	var T float64 = D / S
	fmt.Println(float64(B) + T)

	D = length * (116 - 111) / 360
	T = D / S
	fmt.Println(float64(B) + T)
}
