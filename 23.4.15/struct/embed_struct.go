package main

import "fmt"

type ming struct {
	girl bool
}

type person struct {
	ming
}

func createMing() {
	m := person{ming{
		girl: true,
	}}

	// 匿名导入特性
	fmt.Println(m.girl)
	fmt.Printf("%v", m.ming.girl)
}
