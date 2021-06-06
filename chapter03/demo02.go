package main

import "fmt"

//struct就是一组字段field
type Vertex struct {
	x int
	y int
}

var (
	v1 = Vertex{1, 2}
	v2 = Vertex{x: 1}//Y:0 被隐式地赋予
	v3 = Vertex{}
	v4 = &Vertex{1, 2}
)

func main() {
	fmt.Println(Vertex{x: 1, y: 2})
	fmt.Println(Vertex{1, 2})

	//创建一个结构体变量
	v := Vertex{1, 2}
	v.x = 2
	fmt.Println(v)


	//创建结构体的指针
	pv := &Vertex{1, 2}
	fmt.Println(*pv, pv.x)

	fmt.Println(v1, v2, v3, *v4)
}
