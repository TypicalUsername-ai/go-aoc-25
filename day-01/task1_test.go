package main

import (
	"testing"
)

func TestTask1Case1(t *testing.T) {
	testData := []string{
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

	task := task1{}

	out, h_err := harnessSlice(testData, &task)
	if h_err != nil {
		t.Errorf("harness error: %v", h_err)
	}

	result := make([]byte, 1)
	_, err := out.Read(result)
	if err != nil {
		t.Errorf("an error ocurred %v", err)
	}
	if string(result) != "3" {
		t.Errorf("Task 1: got %s expected 3", result)
	}
}

func TestTask1Case2(t *testing.T) {

	task := task1{}

	out, h_err := harnessFile("inputs/day1.txt", &task)

	if h_err != nil {
		t.Errorf("harness error: %v", h_err)
	}

	result := make([]byte, 4)
	_, err := out.Read(result)
	if err != nil {
		t.Errorf("an error ocurred %v", err)
	}
	if string(result) != "1102" {
		t.Errorf("Task 1: got %s expected 1102", result)
	}
}
