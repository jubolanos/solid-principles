package after

import (
	"encoding/csv"
	"fmt"
	"os"
	"strings"
)

type CSVFileProcessor struct {
}

func NewCSVFileProcessor() *CSVFileProcessor {
	return &CSVFileProcessor{}
}

func (fp CSVFileProcessor) ReadCSVData(filename string) ([][]string, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	reader := csv.NewReader(file)
	data, err := reader.ReadAll()
	if err != nil {
		return nil, err
	}
	return data, nil
}

func (fp CSVFileProcessor) ProcessCSVData(data [][]string) []string {
	var processedData []string
	for _, record := range data {
		processedData = append(processedData, strings.ToUpper(strings.Join(record, ",")))
	}
	fmt.Printf("> Processed Data: %s \n", processedData)
	return processedData
}
