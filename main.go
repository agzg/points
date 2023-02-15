package main

import (
	"fmt"
	"math"
)

type Point struct {
	X, Y int
}

func (p Point) String() string {
	return fmt.Sprintf("(%d, %d)", p.X, p.Y)
}

func (p Point) Distance(q Point) float64 {
	x := math.Pow(float64(p.X-q.X), 2)
	y := math.Pow(float64(p.Y-q.Y), 2)
	return math.Sqrt(x + y)
}

func main() {
	p := Point{0, 3}
	q := Point{4, 0}

	fmt.Println(p, q)
	fmt.Println(p.Distance(q))
}
