package main

import "fmt"

func main() {
	// switch 作用域？
	switch s := "erich"; s {
	case "eric":
		fmt.Println("123")
		break
	default:
		panic(fmt.Sprintf("invalid case, %s", s))
	}
}
