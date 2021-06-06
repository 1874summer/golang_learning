package main

import "fmt"

var (
	a = 10
)

func sumOfNum(num int) (res int) {
	//for循环1
	//for i := 0; i < num; i++ {
	//	res += i
	//}

	//for循环2
	//i := 0
	//for ; i < num; {
	//	res += i
	//	i++
	//}

	//for循环3，相当于while
	//i := 0
	//for i < num {
	//	res += i
	//	i++
	//}

	//for循环4，无限循环
	i := 0
	for {
		res += i
		i++
		if i == num {
			break
		}
	}
	return
}

func main() {
	fmt.Println(sumOfNum(a))
}
