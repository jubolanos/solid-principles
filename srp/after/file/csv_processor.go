package file

import (
	"encoding/csv"
	"io"
	"os"
	"strings"
)

type CSVProcessor struct{}

func NewCSVProcessor() *CSVProcessor {
	return &CSVProcessor{}
}

func (cp CSVProcessor) ReadAndProcessCSV(filename string) ([]string, error) {
	var processedData []string

	file, err := os.Open(filename)

	if err != nil {
		return processedData, err
	}

	defer file.Close()

	reader := csv.NewReader(file)

	for {
		record, err := reader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			return processedData, err
		}

		processedData = append(processedData, strings.ToUpper(strings.Join(record, ",")))
	}

	return processedData, nil
}
