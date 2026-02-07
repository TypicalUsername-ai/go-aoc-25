package main

import (
	"bufio"
	"io"
	"os"
)

func harnessSlice(data []string, task *AocTask) (*bufio.Reader, error) {
	inRead, inWrite := io.Pipe()
	outRead, outWrite := io.Pipe()

	dataIn := bufio.NewWriter(inWrite)
	for _, entry := range data {
		dataIn.WriteString(entry)
	}
	_, taskErr := (*task).Execute(bufio.NewReader(inRead), bufio.NewWriter(outWrite))
	if taskErr != nil {
		return nil, taskErr
	}
	return bufio.NewReader(outRead), nil
}

func harnessFile(filename string, task *AocTask) (*bufio.Reader, error) {
	file, file_err := os.Open(filename)
	if file_err != nil {
		return nil, file_err
	}
	defer file.Close()

	input := bufio.NewReader(file)

	resOut, resIn := io.Pipe()

	_, taskErr := (*task).Execute(input, bufio.NewWriter(resIn))
	if taskErr != nil {
		return nil, taskErr
	}

	return bufio.NewReader(resOut), nil
}
