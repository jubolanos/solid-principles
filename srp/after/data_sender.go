package after

import (
	"fmt"
	"net/http"
	"strings"
)

type DataSender interface {
	SendData(apiUrl string, data []string) error
}

type APISender struct {
}

func NewAPISender() *APISender {
	return &APISender{}
}

func (as *APISender) SendData(apiUrl string, data []string) error {
	for _, processedData := range data {
		fmt.Printf("> Sending data to: %s \n", apiUrl)
		fmt.Printf("> Data to be Sent: %s \n", processedData) // Imprimir datos a enviar
		_, err := http.Post(apiUrl, "text/plain", strings.NewReader(processedData))
		if err != nil {
			return err
		}
	}
	return nil
}
