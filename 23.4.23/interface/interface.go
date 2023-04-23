package main

import "fmt"

type IPerson interface {
	add(s string) string
}

type Person struct {
	A string
}

func (p Person) add(s string) string {
	return s + "1"
}

func main() {
	var p IPerson
	p = Person{}

	// 打印不出来
	// fmt.Println(p.A)

	fmt.Println(p.add("B"))
}
