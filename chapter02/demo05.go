package main

import (
	"fmt"
)

var S1 =1
func fun1() string {
	fmt.Println("fun1")
	return "fun1"
}

func main() {
	defer fmt.Println(fun1())

	fmt.Println("hello")

	fmt.Println("counting")

	for i := 0; i < 10; i++ {
		defer fmt.Println(i)
	}

	fmt.Println("done")
}

//defer 语句会将函数推迟到外层函数返回之后执行
//推迟调用的函数其参数会立即求值，但直到外层函数返回前该函数都不会被调用

//参数会先进行计算，但是函数最后才调用
