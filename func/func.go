package main

import (
	"fmt"
	"math"
)

type Point struct {
	X, Y float64
}

func (p Point) Distance(q Point) float64 {
	return math.Hypot(q.X-q.Y, q.Y-p.Y)
}

func main() {
	p := Point{1, 2}
	q := Point{4, 6}

	c := map[string]int{}
	fmt.Println(c)

	fmt.Println(p.Distance(q))
}
