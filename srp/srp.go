package srp

import (
	"fmt"

	"github.com/jubolanos/solid-principles/srp/after"
)

func Run() {
	fmt.Print("==============================\n")
	fmt.Print("SRP Running... \n")

	fileName := "./resources/cities.csv"
	finalUrl := "http://google.com.mx"
	var err error

	// Before refactor
	// fileProcessor := before.NewFileProcessor()
	// err := fileProcessor.ReadProcessAndSendData(fileName, finalUrl)
	// if err != nil {
	//	fmt.Printf("Ocurrion un error al procesar el archivo. %s", err)
	//}
	// Put your solution here
	csvFileProcessor := after.NewCSVFileProcessor()
	data, err := csvFileProcessor.ReadCSVData(fileName)
	if err != nil {
		fmt.Printf("Ocurrion un error al procesar el archivo. %s", err)
	}
	processedData := csvFileProcessor.ProcessCSVData(data)
	dataSender := after.NewDataSender()
	err = dataSender.SendDataToAPI(finalUrl, processedData)
	if err != nil {
		fmt.Printf("Ocurrion un error al enviar la data. %s", err)
	}

}
