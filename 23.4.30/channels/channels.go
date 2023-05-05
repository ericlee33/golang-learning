package main

import "fmt"

func main() {
	ch1 := make(chan int)
	ch2 := make(chan int)

	go func() {
		ch1 <- 2

	}()

	go func() {
		f := <-ch1
		ch2 <- f

	}()

	fmt.Println(<-ch2)
}
