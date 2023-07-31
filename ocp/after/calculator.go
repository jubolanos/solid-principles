package after

type Calculator struct {
}

func (c Calculator) SumAllAreas(shapes ...Shape) float32 {
	var sum float32
	for _, shape := range shapes {
		sum = shape.CalculateArea()
	}
	return sum
}
