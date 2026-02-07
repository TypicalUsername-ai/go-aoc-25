package main

import (
	"bufio"
	"log"
	"strconv"
	"strings"
)

type Task1 struct{}

func (t *Task1) Execute(input *bufio.Reader, output *bufio.Writer) (int, error) {

	//reading directions
	var dialDirections []string
	for {
		line, err := input.ReadString('\n')
		if err != nil {
			log.Fatal(err)
			return -1, err
		}
		data := strings.TrimSpace(line)
		if len(data) == 0 {
			break
		}
		dialDirections = append(dialDirections, data)
	}
	// parsing directions
	dialPosition := 50
	ticker := 0

	for _, dir := range dialDirections {
		num, err := strconv.Atoi(dir[1:])
		if err != nil {
			log.Fatal(err)
			return -1, err
		}
		if dir[0] == 'L' {
			dialPosition -= num
		} else {
			dialPosition += num
		}
		if dialPosition%100 == 0 {
			ticker += 1
		}

	}
	output.WriteString(strconv.FormatInt(int64(ticker), 10))

	return 1, nil

}
