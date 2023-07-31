package after

import (
	"fmt"
	"net/http"
	"strings"
)

type DataSender struct {
}

func NewDataSender() *DataSender {
	return &DataSender{}
}

func (ds DataSender) SendDataToAPI(apiUrl string, data []string) error {
	fmt.Printf("> Sending data to: %s \n", apiUrl)
	for _, record := range data {
		_, err := http.Post(apiUrl, "text/plain", strings.NewReader(record))
		if err != nil {
			return err
		}
	}
	return nil
}
