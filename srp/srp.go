package srp

import (
	"fmt"
	"github.com/jubolanos/solid-principles/srp/after"

	"github.com/jubolanos/solid-principles/srp/before"
)

func Run() {
	fmt.Print("SRP Running... \n")

	// Before refactor
	fileProcessor := before.NewFileProcessor()
	err := fileProcessor.ReadProcessAndSendData("./resources/cities.csv", "http://google.com.mx")
	if err != nil {
		fmt.Printf("Ocurrion un error al procesar el archivo. %s", err)
	}

	// Put your solution here
	reader := after.NewCvsReader()
	processor := after.NewFileProcessor()
	sender := after.NewAPISender()

	data, err := reader.ReadFile("./resources/cities.csv")
	if err != nil {
		fmt.Printf("Ocurrion un error al leer el archivo. %s", err)
	}

	processedData := processor.ProcessFile(data)
	sender.SendData("http://google.com.mx", processedData)
}
