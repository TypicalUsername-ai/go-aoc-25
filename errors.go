package main

import (
	"fmt"
)

type NoSuchTask struct {
	num int
}

func (e *NoSuchTask) Error() string {
	if e.num > 24 || e.num < 1 {

		return fmt.Sprintf("Invalid task number %v, please provide one between 1 and 24", e.num)
	} else {

		return fmt.Sprintf("Task number %v has not been implemented yet", e.num)
	}
}
