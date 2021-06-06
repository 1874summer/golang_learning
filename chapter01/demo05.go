package main

import (
	"fmt"
	"math"
)

//go的基本类型
/*
数值型：
int int8 int16 int32 int64
uint uint8 uint16 uint32 uint64 uintptr
byte //相当于uint8
float32 float64
bool
string
*/
var z int = 1
var (
	a1 = 1
	d1 = 1.00
	b1 = true
	c1 string
)

func main() {
	fmt.Printf("%T,%t,%T,%t,\n", a1, b1, c1, d1)
	fmt.Printf("%T,%v,\n",a1,a1)
	fmt.Printf("%T,%q",c1,c1)//string的值得用"%q"
	num := 64
	fmt.Println(math.Sqrt(float64(num)))
}
