package after

import "github.com/jubolanos/solid-principles/ocp/after/shapes"

type Calculator struct{}

func (c Calculator) SumAllAreas(shapes ...shapes.Shape) float32 {
	var sum float32

	for _, shape := range shapes {
		sum += shape.Area()
	}

	return sum
}
