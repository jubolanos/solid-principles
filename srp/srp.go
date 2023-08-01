package srp

import (
	"fmt"

	"github.com/jubolanos/solid-principles/srp/after"
	"github.com/jubolanos/solid-principles/srp/before"
)

func Run() {
	fmt.Print("SRP Running... \n")
	Before()
	After()
}

func Before() {
	// Before refactor
	fileProcessor := before.NewFileProcessor()
	err := fileProcessor.ReadProcessAndSendData("./resources/cities.csv", "http://google.com.mx")
	if err != nil {
		fmt.Printf("Ocurrion un error al procesar el archivo. %s", err)
	}

}

func After() {
	// After refactor
	apiUrl := "http://google.com.mx"
	filename := "./resources/cities.csv"

	fileProcessor := after.NewFileProcessor(apiUrl)
	err := fileProcessor.ProcessFileAndSendData(filename)

	if err != nil {
		fmt.Println(err)
		return
	}
}
