package shapes

type Circle struct {
	Radius float32
}

func (c Circle) Area() float32 {
	return c.Radius * c.Radius * 3.14159
}
