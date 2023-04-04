package main

import "fmt"

func counter(out chan<- int) {
	for i := 0; i < 100; i++ {
		out <- i
	}
	// 如果没有 close，就会 deadlock 死锁，因为 squarer 的方法是 range，会继续阻塞等待
	close(out)
}

func squarer(out chan<- int, in <-chan int) {
	for v := range in {
		out <- v * v
	}
	close(out)
}

func pointer(in <-chan int) {
	for v := range in {
		fmt.Println((v))
	}
}

func main() {
	naturals := make(chan int)
	squares := make(chan int)

	go counter(naturals)
	go squarer(squares, naturals)

	pointer(squares)
}
