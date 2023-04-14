package main

import "fmt"

func main() {
	// 默认空 数组
	var a [3]int

	fmt.Println(a[0])

	// var b [3]int = [3]int{1, 2, 3}
	// 或者
	b := [...]int{1, 2, 3}

	for _, v := range b {
		fmt.Println(v)
	}

}
