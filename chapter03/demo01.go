package main

import "fmt"
//var S1 = 1
func main() {

	i := 1
	//申明一个指针
	var p1 *int
	//给指针p1赋值
	p1 = &i
	//通过指针p1访问变量i
	fmt.Println(*p1)

	j := 2
	p2 := &j

	fmt.Println(p1, p2)
	fmt.Println(*p1, *p2)
}
