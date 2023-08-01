package before

import "fmt"

// Electric Vehicle
type ElectricVehicle struct {
	Vehicle
}

func (ev ElectricVehicle) PrintInfo() {
	fmt.Printf("ElectricVehicle's manufacturing year is: %d\n", ev.ManufacturingYear)
}

func (ev ElectricVehicle) Emissions() int {
	return 0
}
