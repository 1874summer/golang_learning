package main

import "fmt"

var (
	a6 = 1
	b6 = 2
)

const (
	A = 1 << 100
	B = 1
)

func nextInt(num int) (res int) {
	res = num
	return
}

func nextFloat(num float64)  float64 {
	return num
}

func main() {
	//一个未指定类型的常量由上下文来决定其类型
	fmt.Println(nextInt(B))
	fmt.Println(nextFloat(A))
}
