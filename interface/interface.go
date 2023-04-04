package main

import (
	"fmt"
)

// declare a rectangle struct
type rectangle struct {
	length int
	width  int
}

// declare an interface with area() as a member
type shape interface {
	area() int
}

// declare a method area()
// the rectangle struct implements area() method in shape interface
func (r rectangle) area() int {
	return r.length * r.width
}

// declare a method with type shape as a parameter
func info(s shape) {
	fmt.Println("the area: ", s.area())
}

func main() {
	r1 := rectangle{12, 12}
	//r1 is a rectangle type. rectangle implements all methods in shape interface.
	info(r1)
}

// output
// the area:  144
