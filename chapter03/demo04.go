package main

import "fmt"

func main() {

	//切片
	//每个数组的大小都是固定的。而切片则为数组元素提供动态大小的、灵活的视角。
	a := [2]int{1, 2}
	fmt.Println(a)

	a1 := a[1:2]
	a1[0] = 3
	fmt.Println(a1, a1[0])
	fmt.Println(a[1]) //改变slice的值，也会改变原数组的值
	//切片就像数组的引用，更改切片的元素会修改其底层数组中对应的元素

	fmt.Printf("%T,%T\n", a, a1) //[2]int,[]int

	i := []int{1, 2, 3}
	j := []bool{true, false}
	k := []struct {
		num1 int
		num2 int
	}{
		{1, 2},
		{1, 3},
	}
	fmt.Println(i, j, k)//[1 2 3] [true false] [{1 2} {1 3}]
	fmt.Printf("%T,%T,%T\n", i, j, k)//[]int,[]bool,[]struct { num1 int; num2 int }

}
