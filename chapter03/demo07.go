package main

import (
	"fmt"
)

//用 make 创建切片
func main() {

	//make([]T, a, b) 创建一个数组[]T，并且返回其一个slice引用，a为len，b为cap
	a := make([]int, 2, 10)
	fmt.Println(len(a), cap(a))

	b := make([]int, 2)//长度为2，cap也为2
	fmt.Println(len(b), cap(b))

}
func printSlice1(s string, x []int) {
	fmt.Printf("%s len=%d cap=%d %v\n",
		s, len(x), cap(x), x)

}
