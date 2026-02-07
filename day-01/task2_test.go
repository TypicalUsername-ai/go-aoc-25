package main

import (
	"testing"
)

func TestTask2Case1(t *testing.T) {
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
	if string(result) != "6" {
		t.Errorf("Task 2: got %s expected 6", result)
	}
}

func TestTask2Case2(t *testing.T) {
	testData := []string{
		"L200", // ticks 2 times -> 50
		"R350", // ticks 4 times -> 0
		"L55",  // ticks 0 times -> 55
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
	if string(result) != "6" {
		t.Errorf("Task 2: got %s expected 6", result)
	}
}

func TestTask2Case3(t *testing.T) {
	t.Skip()
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
	if string(result) != "6" {
		t.Errorf("Task 2: got %s expected 1102", result)
	}
}
