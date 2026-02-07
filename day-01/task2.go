package main

import (
	"bufio"
	"fmt"
	"log"
	"strconv"
	"strings"
)

type task2 struct {
}

func (t *task2) Execute(input *bufio.Reader, output *bufio.Writer) (int, error) {
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
		fmt.Printf("%v\n ", dialPosition)
		num, err := strconv.Atoi(dir[1:])
		if err != nil {
			log.Fatal(err)
		}
		if num >= 100 {
			//fmt.Printf("\t+%v %s->%c%v\n", num/100, dir, dir[0], num%100)
			ticker += num / 100
			num = num % 100
		}
		if dir[0] == 'L' {
			if num > dialPosition {
				//	fmt.Printf("\tn:%v p:%v\n", num, dialPosition)
				dialPosition += 100
				if dialPosition != 0 {
					ticker += 1
				}
				//fmt.Printf("\t+1\n")
			}
			dialPosition -= num
		} else {
			dialPosition += num
			if dialPosition > 99 {
				dialPosition -= 100
				ticker += 1
				fmt.Printf("\t+1\n")
			}
		}
		if dialPosition == 0 {
			ticker += 1
			//fmt.Printf("\t+1\n")
		}

	}
	//fmt.Printf("[AoC] %v - %v\n", dialPosition, ticker)
	output.WriteString(strconv.FormatInt(int64(ticker), 10))

	return 1, nil

}
