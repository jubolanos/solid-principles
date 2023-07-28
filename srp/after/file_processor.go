package after

import (
	"fmt"
	"strings"
)

type FileProcessor interface {
	ProcessFile([]string) []string
}

type CvsProcessor struct{}

func NewFileProcessor() FileProcessor {
	return &CvsProcessor{}
}

func (fp *CvsProcessor) ProcessFile(data []string) []string {
	var processedData []string
	for _, record := range data {
		processedData = append(processedData, strings.ToUpper(record))
		fmt.Printf("> Processed Data: %s \n", strings.ToUpper(record)) // Imprimir datos procesados
	}
	return processedData
}
