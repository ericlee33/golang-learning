package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan int) // 用来通信用的管道
	// 通过 「go 关键字 + 执行函数」 开启一个协程
	go func(ch chan int) {
		time.Sleep(time.Duration(3) * time.Second)
		ch <- 1 // 向管道中塞入 1.
	}(ch)
	v := <-ch // 从 管道中接收数据
	fmt.Println(v)
	fmt.Println(2)
}
