package before

type FileProcessor struct {
}

func NewFileProcessor() *FileProcessor {
	return &FileProcessor{}
}

func (fp FileProcessor) ReadProcessAndSendData(filename string, apiUrl string) error {
	//for {
	//	record, err := reader.Read()
	//	if err == io.EOF {
	//		break
	//	}
	//	if err != nil {
	//		return err
	//	}
	//	processedData := strings.ToUpper(strings.Join(record, ","))
	//	fmt.Printf("> Processed Data: %s \n", processedData)
	//
	//	fmt.Printf("> Sending data to: %s \n", apiUrl)
	//	_, err = http.Post(apiUrl, "text/plain", strings.NewReader(processedData))
	//	if err != nil {
	//		return err
	//	}
	//}
	//return nil
	return nil
}
