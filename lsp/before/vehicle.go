package before

import "fmt"

// Vehicle
type Vehicle struct {
	ManufacturingYear int
}

func (v Vehicle) PrintInfo() {
	fmt.Printf("Vehicle's manufacturing year is: %d\n", v.ManufacturingYear)
}

func (v Vehicle) Emissions() int {
	return 200
}

func (v Vehicle) PrintEmissions() {
	fmt.Printf("Vehicle's emissions: %d\n", v.Emissions())
}
