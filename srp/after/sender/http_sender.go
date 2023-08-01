package sender

import (
	"fmt"
	"net/http"
	"strings"
)

type DataSender struct {
	APIUrl string
}

func NewDataSender(apiUrl string) *DataSender {
	return &DataSender{APIUrl: apiUrl}
}

func (ds DataSender) SendData(data []string) error {

	for _, processedData := range data {
		fmt.Printf("> Processed Data: %s \n", processedData)
		fmt.Printf("> Sending data to: %s \n", ds.APIUrl)

		_, err := http.Post(ds.APIUrl, "text/plain", strings.NewReader(processedData))

		if err != nil {
			return err
		}
	}

	return nil
}
