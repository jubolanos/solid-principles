package before

type Calculator struct {
}

func (c Calculator) SumAllAreas(shapes ...Shape) float64 {
	var sum float64
	for _, shape := range shapes {
		sum = shape.CalculateArea()
	}
	return sum
}
