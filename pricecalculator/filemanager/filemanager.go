package filemanager

import (
	"bufio"
	"encoding/json"
	"errors"
	"os"
	"time"
)

type FileManager struct {
	InputFileName  string
	OutPutFileName string
}

func (fm FileManager) ReadLines() ([]string, error) {
	file, err := os.Open(fm.InputFileName)

	if err != nil {
		return nil, errors.New("failed to open file")
	}

	scanner := bufio.NewScanner(file)

	var lines []string

	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	err = scanner.Err()
	if err != nil {
		file.Close()
		return nil, errors.New("failed to read lines")
	}

	file.Close()
	return lines, nil
}

func (fm FileManager) WriteResult(data any) error {
	file, err := os.Create(fm.OutPutFileName)

	if err != nil {
		return errors.New("failed to write result")
	}

	time.Sleep(3 * time.Second)

	encoder := json.NewEncoder(file)
	err = encoder.Encode(data)

	if err != nil {
		file.Close()
		return errors.New("failed to encode result")
	}

	file.Close()
	return nil
}

func New(inputfile string, outputfile string) FileManager {
	return FileManager{
		InputFileName:  inputfile,
		OutPutFileName: outputfile,
	}
}
