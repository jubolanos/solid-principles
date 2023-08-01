package shapes

type Triangle struct {
	Base   float32
	Height float32
}

func (t Triangle) Area() float32 {
	return t.Base * t.Height / 2
}
