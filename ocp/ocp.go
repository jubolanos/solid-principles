package ocp

import (
	"fmt"

	"github.com/jubolanos/solid-principles/ocp/after"
	"github.com/jubolanos/solid-principles/ocp/after/shapes"
	"github.com/jubolanos/solid-principles/ocp/before"
)

func Run() {
	Before()
	After()
}

func Before() {

	fmt.Println("Open Closed Principle...")
	circle := before.Circle{Radius: 5}
	square := before.Square{Length: 10}
	calculator := before.Calculator{}
	totalSum := calculator.SumAllAreas(circle, square)

	fmt.Println("The total sum is: ", totalSum)

}

func After() {

	fmt.Println("Open Closed Principle...")

	circle := shapes.Circle{Radius: 5}
	square := shapes.Square{Length: 10}

	//triangle := shapes.Triangle{Base: 10, Height: 5}
	//rectangle := shapes.Rectangle{Width: 10, Height: 5}

	calculator := after.Calculator{}
	//totalSum := calculator.SumAllAreas(circle, square, triangle, rectangle)
	totalSum := calculator.SumAllAreas(circle, square)

	fmt.Println("The total sum is: ", totalSum)
}
