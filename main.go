package main

import (
	"fmt"
	"log"
	"strconv"
)

type AocChallenge interface {
	GetInput() AocChallenge
	Challenge() string
}

func main() {
	fmt.Println("which day task you want to run?:")
	var input string
	fmt.Scan(&input)
	num, err := strconv.Atoi(input)
	if err != nil {
		log.Fatal(err)
	}
	var day AocChallenge
	switch {
	case num == 1:
		day = day1{data: []string{}}
	default:
		log.Fatal("could not find task")
	}
	day = day.GetInput()
	output := day.Challenge()
	fmt.Printf("[result] %v\n", output)
}
