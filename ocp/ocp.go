package ocp

import (
	"fmt"
	after "github.com/jubolanos/solid-principles/ocp/after"

	"github.com/jubolanos/solid-principles/ocp/before"
)

func Run() {

	fmt.Println("Open Closed Principle...")
	beforeCircle := before.Circle{Radius: 5}
	beforeSquare := before.Square{Length: 10}
	beforeCalculator := before.Calculator{}
	beforeTotalSum := beforeCalculator.SumAllAreas(&beforeCircle, &beforeSquare)

	afterCircle := after.Circle{Radius: 5}
	afterSquare := after.Square{Length: 10}
	afterCalculator := after.Calculator{}
	afterTotalSum := afterCalculator.SumAllAreas(&afterCircle, &afterSquare)

	fmt.Println("BEFORE: The total sum is: ", beforeTotalSum)
	fmt.Println("AFTAR: The total sum is: ", afterTotalSum)

}
