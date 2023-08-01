package ocp

import (
	"fmt"

	"github.com/jubolanos/solid-principles/ocp/after"
)

func Run() {

	fmt.Print("==============================\n")
	fmt.Println("Open Closed Principle...")
	// circle := before.Circle{Radius: 5}
	//square := before.Square{Length: 10}
	// calculator := before.Calculator{}
	// totalSum := calculator.SumAllAreas(circle, square)

	circle := after.Circle{Radius: 10}
	square := after.Square{Length: 8}
	triangle := after.Triangle{Base: 12, Height: 6}

	ocpCalculator := after.Calculator{}

	totalSum := ocpCalculator.SumAllAreas(circle, square, triangle)

	fmt.Println("The total sum is: ", totalSum)

}
