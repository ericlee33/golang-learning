// package main

// import (
// 	"fmt"
// 	"sync"
// 	"time"
// )

// var m sync.Mutex

// func main() {
// 	var arr []int

// 	for i := 0; i < 5; i++ {
// 		go setValue(&arr, i)
// 	}

// 	time.Sleep(time.Second * 3)
// 	fmt.Println(arr)
// }

// func setValue(arr *[]int, n int) {
// 	m.Lock()

// 	*arr = append(*arr, n*n)
// 	m.Unlock()
// }

package main

import (
	"fmt"
	"sync"
	"time"
)

var m sync.Mutex

func main() {
	var arr []int

	for i := 0; i < 5; i++ {
		go setValue(&arr, i)
	}

	time.Sleep(time.Second * 3)
	fmt.Println(arr)
}

func setValue(arr *[]int, n int) {
	m.Lock()
	fmt.Println(n)
	defer m.Unlock()

	*arr = append(*arr, n*n)
}
