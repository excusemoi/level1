package task24

import (
	"fmt"
	"math"
)

type Point struct {
	x float64
	y float64
}

func NewPoint(x, y float64) *Point{
	return &Point{
		x: x,
		y: y,
	}
}

func Distance(p1, p2 *Point) float64 {
	return math.Sqrt(math.Pow(p1.x-p2.x, 2) + math.Pow(p1.y-p2.y, 2))
}

func Solve(p1, p2 *Point){
	fmt.Printf("Distance between points %v and %v:\n%f\n", *p1, *p2, Distance(
		p1,
		p2,
	))
}