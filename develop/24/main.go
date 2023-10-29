package main

import (
	"fmt"
	"math"
)

type Point struct {
	x float64
	y float64
}

func (p Point) Distance(s Point) float64 {
	length := math.Sqrt(math.Pow((p.x - s.x), 2) + math.Pow((p.y - s.y), 2))
	return length
}

func New(x, y float64) Point {
	return Point{
		x: x,
		y: y,
	}
}

func main() {
	first := New(3, 2)
	second := New(1, 2)

	fmt.Printf("%v", first.Distance(second))
}
