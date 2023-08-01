package srp

import (
	"fmt"

	"github.com/jubolanos/solid-principles/srp/before"
)

func Run() {
	fmt.Print("==============================\n")
	fmt.Print("SRP Running... \n")

	// Before refactor
	fileProcessor := before.NewFileProcessor()
	err := fileProcessor.ReadProcessAndSendData("./resources/cities.csv", "http://google.com.mx")
	if err != nil {
		fmt.Printf("Ocurrion un error al procesar el archivo. %s", err)
	}

	// Put your solution here
}
