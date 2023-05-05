package main

import (
	"fmt"
	"time"
)

func main() {
	go spin()

	res := fib(45)

	fmt.Println(res)
}

func spin() {
	for {
		fmt.Printf("\r-")
		time.Sleep(1000 * time.Millisecond)
	}
}

func fib(n int) int {
	if n < 2 {
		return n
	}
	return fib(n-1) + fib(n-2)
}

// func main() {
// 	go spinner(100 * time.Millisecond)
// 	const n = 45
// 	fibN := fib(n) // slow
// 	fmt.Printf("\rFibonacci(%d) = %d\n", n, fibN)
// }

// func spinner(delay time.Duration) {
// 	for {
// 		for _, r := range `-\|/` {
// 			fmt.Printf("\r%c", r)
// time.Sleep(delay)
// 		}
// 	}
// }

// func fib(x int) int {
// 	if x < 2 {
// 		return x
// 	}
// 	return fib(x-1) + fib(x-2)
// }
