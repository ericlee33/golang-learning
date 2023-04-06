package main

import "fmt"

func main() {
	a := test()

	*a = 2
	fmt.Println(*a)
}

func test() *int {
	a := 1
	return &a
}
