package shapes

type Square struct {
	Length float32
}

func (s Square) Area() float32 {
	return s.Length * s.Length
}
