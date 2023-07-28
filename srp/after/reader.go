package after

import (
	"encoding/csv"
	"io"
	"os"
)

type Reader interface {
	ReadFile(filename string) ([]string, error)
}

type CvsReader struct{}

func NewCvsReader() *CvsReader {
	return &CvsReader{}
}

func (cr *CvsReader) ReadFile(filename string) ([]string, error) {
	var data []string
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	reader := csv.NewReader(file)
	for {
		record, err := reader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			return nil, err
		}
		data = append(data, record...)
	}
	return data, nil
}
