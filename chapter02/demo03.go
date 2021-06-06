package main

import (
	"fmt"
	"math"
)

func pow(x float64, n float64, limit float64) float64 {
	if temp := math.Pow(x, n); temp < limit {
		return temp
	}
	return limit
}

func main() {
	fmt.Println(pow(2,10,100))
	fmt.Println(pow(2,4,100))
}
