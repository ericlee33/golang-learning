package main

import (
	"fmt"
)

func main() {

	// 构建一个通道
	ch := make(chan int)

	// 开启一个并发匿名函数
	go func() {
		// 通过通道通知main的goroutine

		fmt.Println("start goroutine")

		for i := 0; i < 10000; i++ {
			fmt.Println("exit %d goroutine", i)

		}
		ch <- 0

	}()

	fmt.Println("wait goroutine")

	// 等待匿名goroutine
	// for i := 0; i < 2; i++ {
	// }
	<-ch

	fmt.Println("all done")

}
