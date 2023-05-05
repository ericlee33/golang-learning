package main

import "fmt"

func main() {
	c := make(chan int)

	go set(c)

	d := <-c

	fmt.Println(d)
}

func set(in chan<- int) {
	in <- 3
}
