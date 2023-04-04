package main

import "fmt"

func main() {
	a := 1
	b := &a
	c := *b
	fmt.Println("hello, world", c, c)

	fmt.Print("hello, world%d-", c)

	fmt.Fprintf()
}

// func abc() {
// 	// b := 2
// 	c := '1'

// 	fmt.Println("hello%s", c)

// 	return
// }
