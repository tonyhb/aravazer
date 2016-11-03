package main

import (
	"fmt"
	"math"
)

type Point struct {
	X int
	Y int
}

func (p Point) Expand() []Point {
	return []Point{
		{p.X - 1, p.Y},
		{p.X + 1, p.Y},
		{p.X, p.Y - 1},
		{p.X, p.Y + 1},
	}
}

func (p Point) ToString() string {
	return fmt.Sprintf("%d,%d", p.X, p.Y)
}

func (p Point) ManhattanDistance(from Point) int {
	return int(math.Abs(float64(p.X-from.X)) + math.Abs(float64(p.Y-from.Y)))
}
