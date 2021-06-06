package main

import "fmt"

func printSlice(s []int) {
	fmt.Printf("长度为：%d,容量为%d。\n", len(s), cap(s))
}

func main() {
	//切片的长度和容量

	s := []int{2, 3, 5, 7, 11, 13}
	printSlice(s)

	// 截取切片使其长度为 0
	s = s[:0]
	printSlice(s)

	// 拓展其长度
	s = s[:4]
	printSlice(s)

	// 舍弃前两个值
	s = s[2:]
	printSlice(s)


	//nil切片的len和cap都为0，且没有底层数组
	var j []int
	fmt.Println(j, len(j), cap(j))
	if j == nil {
		fmt.Println("nil!")
	}

}
