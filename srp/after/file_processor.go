package after

import (
	"fmt"

	"github.com/jubolanos/solid-principles/srp/after/file"
	"github.com/jubolanos/solid-principles/srp/after/sender"
)

type FileProcessor struct {
	csvProcessor *file.CSVProcessor
	dataSender   *sender.DataSender
}

func NewFileProcessor(apiUrl string) *FileProcessor {
	return &FileProcessor{
		csvProcessor: file.NewCSVProcessor(),
		dataSender:   sender.NewDataSender(apiUrl),
	}
}

func (fp *FileProcessor) ProcessFileAndSendData(filename string) error {
	processedData, err := fp.csvProcessor.ReadAndProcessCSV(filename)
	if err != nil {
		return fmt.Errorf("error processing CSV: %v", err)
	}

	err = fp.dataSender.SendData(processedData)
	if err != nil {
		return fmt.Errorf("error sending data to API: %v", err)
	}

	return nil
}
