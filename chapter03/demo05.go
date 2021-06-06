package main

import "fmt"

func main() {
	s := [3]int{1, 2, 3}
	q := [...]int{1, 2, 3}

	/*
	以下切片是等价的：
	a[0:10]
	a[:10]
	a[0:]
	a[:]
	 */
	s1 := s[0:3]
	fmt.Println(s1)

	s2 := s[0:]
	fmt.Println(s2)

	s3 := s[:3]
	fmt.Println(s3)

	s4 := s[:]
	fmt.Println(s4)

	//fmt.Println(s1 == s2)//slice 只能与 nil 进行比较
	fmt.Println(s == q)//true ?



}
