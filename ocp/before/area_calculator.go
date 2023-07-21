package before

import (
	"fmt"
	"math"
)

type Calculator struct {
}

func (c Calculator) SumAllAreas(shapes ...interface{}) float32 {
	var sum float32
	for _, shape := range shapes {
		switch v := shape.(type) {
		case Square:
			l := shape.(Square).Length
			sum += l * l
		case Circle:
			r := shape.(Circle).Radius
			sum += math.Pi * r * r
		default:
			fmt.Printf("Unknow type. %T", v)
		}
	}
	return sum
}
