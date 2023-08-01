package shapes

type Rectangle struct {
	Width  float32
	Height float32
}

func (r Rectangle) Area() float32 {
	return r.Width * r.Height
}
