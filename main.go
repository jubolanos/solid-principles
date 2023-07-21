package main

import (
	"fmt"

	"github.com/jubolanos/solid-principles/ocp"
	"github.com/jubolanos/solid-principles/srp"
)

func main() {
	fmt.Println("Learn SOLID principles.")
	srp.Run()
	ocp.Run()

}
