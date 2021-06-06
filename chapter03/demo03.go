package main

import "fmt"

func main() {

	//类型 [n]T 表示拥有 n 个 T 类型的值的数组
	//数组的长度是其类型的一部分，因此数组不能改变大小。
	var a [2]int
	a[0] = 0
	a[1] = 1
	fmt.Println(a[0], a[1])
	fmt.Println(a)

	b := [...]int{1,2}
	fmt.Println(b)
	//b = [3]int{1,2,3}  //报错，b为[2]int类型
	fmt.Printf("%T",b)
}
