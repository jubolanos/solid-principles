package ocp

import (
	"fmt"

	"github.com/jubolanos/solid-principles/ocp/before"
)

func Run() {

	fmt.Print("==============================\n")
	fmt.Println("Open Closed Principle...")
	circle := before.Circle{Radius: 5}

	square := before.Square{Length: 10}

	calculator := before.Calculator{}

	totalSum := calculator.SumAllAreas(circle, square)

	fmt.Println("The total sum is: ", totalSum)

}
