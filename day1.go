package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type day1 struct {
	data []string
}

func (d day1) GetInput() AocChallenge {
	fmt.Println("[AoC day1] Paste your input")
	inputReader := bufio.NewReader(os.Stdin)

	var dialDirections []string
	for {
		line, err := inputReader.ReadString('\n')
		if err != nil {
			log.Fatal(err)
		}
		data := strings.TrimSpace(line)
		if len(data) == 0 {
			break
		}
		dialDirections = append(dialDirections, data)
	}

	d.data = dialDirections
	return d
}

func (d day1) Challenge() string {
	dialPosition := 0
	ticker := 0

	fmt.Println(d.data)

	for _, dir := range d.data {
		num, err := strconv.Atoi(dir[1:])
		if err != nil {
			log.Fatal(err)
		}
		if dir[0] == 'L' {
			dialPosition -= num
		} else {
			dialPosition += num
		}
		fmt.Printf("[AoC 1] %v\n", dialPosition)
		if dialPosition%100 == 0 {
			ticker += 1
		}

	}
	fmt.Printf("[AoC] %v\n", ticker)
	return strconv.FormatInt(int64(ticker), 10)
}
