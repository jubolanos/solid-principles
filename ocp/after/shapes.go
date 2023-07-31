package after

import "math"

type Shape interface {
	CalculateArea() float32
}

// Square def
type Square struct {
	Length float32
}

func (sq Square) CalculateArea() float32 {
	return sq.Length * sq.Length
}

// Circle definition
type Circle struct {
	Radius float32
}

func (c Circle) CalculateArea() float32 {
	return math.Pi * (c.Radius * c.Radius)
}

// Triangle def.

type Triangle struct {
	Base   float32
	Height float32
}

func (t Triangle) CalculateArea() float32 {
	return (t.Base * t.Height) / 2
}
