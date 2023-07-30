package before

import "math"

type Shape interface {
	CalculateArea() float64
}

type Square struct {
	Length float64
}

func (square *Square) CalculateArea() float64 {
	return square.Length * square.Length
}

type Circle struct {
	Radius float64
}

func (circle *Circle) CalculateArea() float64 {
	return math.Pi * circle.Radius * circle.Radius
}
