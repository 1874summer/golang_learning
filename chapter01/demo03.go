package main

import "fmt"

func swap(str1, str2 string) (res1, res2 string) {
	res1 = str2
	res2 = str1
	return
}

func main() {
	str1, str2 := swap("1", "2")
	fmt.Println(str1, str2)
}
