package main

import "fmt"

func main() {
	if a := 1; false {
		// fmt.Println(a, b, c) // undefined: b, c
	} else if b := 2; false {
		// fmt.Println(a, b, c) // undefined: c
	} else if c := 3; false {
		fmt.Println(a, b, c)
	} else {
		fmt.Println(a, b, c)
	}
}
