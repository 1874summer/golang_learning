package main

import "fmt"

//定义变量1
var a, b, c = true, true, true

//定义变量2
var d, e, f bool

func main() {
	//定义变量3
	// := 结构不能在函数外使用
	h, i, j := 1, 2, 3

	var k, l, m int
	k,l,m = 1,2,3
	var d = 1
	fmt.Println(a, b, c, d)
	fmt.Println(h, i, j,k,l,m)
}

//变量定义在方法外，即使不使用也不会报错，定义在方法内的，必须使用否则会报错

//变量的两种定义方式：

//1.var申明，后赋值  或者  var i,j = 1,2
//2.直接用：=表示申明并赋值，但是只能用在方法内
