package main

import (
	"bufio"
	"log"
	"os"
	"strings"
	"testing"
)

func TestDay1Case1(t *testing.T) {
	var input []string
	input = []string{
		"L68",
		"L30",
		"R48",
		"L5",
		"R60",
		"L55",
		"L1",
		"L99",
		"R14",
		"L82",
	}

	day := day1{data: input}
	result := day.Challenge()
	if result != "3" {
		t.Errorf("Day 1: got %s expected 3", result)
	}
}

func TestDay1Case2(t *testing.T) {
	var input []string
	file, err := os.Open("day1.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		input = append(input, strings.TrimSpace(scanner.Text()))
	}

	day := day1{data: input}
	result := day.Challenge()
	if result != "1102" {
		t.Errorf("Day 1: got %s expected 1102", result)
	}
}
