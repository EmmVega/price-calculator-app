package filemanager

import (
	"bufio"
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"time"
)

type Filemanager struct {
	InputFilePath  string
	OutputFilePath string
}

func (fm Filemanager) ReadFiles() ([]string, error) {
	file, err := os.Open(fm.InputFilePath)

	if err != nil {
		fmt.Println("Could not open file!")
		fmt.Println(err)
		return nil, errors.New("failed to open file")
	}

	defer file.Close()

	scaner := bufio.NewScanner(file)

	var lines []string

	for scaner.Scan() {
		lines = append(lines, scaner.Text())
	}

	err = scaner.Err()

	if err != nil {
		//file.Close()
		return nil, errors.New("failed to read file")
	}

	os.Stdin.Close()
	return lines, nil
}

func (fm Filemanager) WriteJSON(data interface{}) error {
	file, err := os.Create(fm.OutputFilePath)

	if err != nil {
		return errors.New("fail to create file")
	}

	defer file.Close()

	time.Sleep(3 * time.Second)

	encoder := json.NewEncoder(file)
	err = encoder.Encode(data)

	if err != nil {
		//file.Close()
		return errors.New("fail to convert data to json")
	}

	return nil
}

func New(inputPath string, outputPath string) Filemanager {
	return Filemanager{
		InputFilePath:  inputPath,
		OutputFilePath: outputPath,
	}
}
