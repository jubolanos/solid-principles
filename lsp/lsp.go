package lsp

import (
	"fmt"

	"github.com/jubolanos/solid-principles/lsp/before"
)

func Run() {
	fmt.Print("==============================\n")
	fmt.Print("LSR Running... \n")

	var vehicle before.Vehicle
	vehicle.ManufacturingYear = 2019

	var electricVehicle before.ElectricVehicle
	electricVehicle.ManufacturingYear = 2020

	// open for extension
	vehicle.PrintInfo()         // Vehicle's manufacturing year is: 2019
	electricVehicle.PrintInfo() // ElectricVehicle's manufacturing year is: 2020

	// closed for modification
	fmt.Println(electricVehicle.Emissions()) // prints 0
	electricVehicle.PrintEmissions()         // prints Vehicle's emissions: 200, instead of 0

	// Dos componentes son sustituibles si exhiben un comportamiento tal que
	// la persona que llama no puede notar la diferencia.

}
