package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
)

type AocTask interface {
	Execute(input *bufio.Reader, output *bufio.Writer) (int, error)
}

func main() {

	taskNo := flag.Int("task", -1, "Task number to execute")
	inputFile := flag.String("in", "/dev/stdin", "Location of input for task data (defaults to stdin)")
	outputFile := flag.String("out", "/dev/stdout", "Where to output task results (default stdout)")

	flag.Parse()

	if *taskNo == -1 {
		fmt.Println("No task number provided, please enter desired task number now:")
		_, err := fmt.Scan(taskNo)
		if err != nil {
			log.Fatal(err)
			panic("Invalid or no task number provided")
		}
	}

	task, err := lookupTask(*taskNo)

	if err != nil {
		log.Fatal(err)
		panic("Task lookup failed")
	}

	inStream, in_err := os.Open(*inputFile)

	outStream, out_err := os.Create(*outputFile)

	var input *bufio.Reader
	var output *bufio.Writer

	if in_err != nil {
		log.Fatal(in_err)
		input = bufio.NewReader(os.Stdin)
	} else {
		input = bufio.NewReader(inStream)
	}

	if out_err != nil {
		log.Fatal(out_err)
		output = bufio.NewWriter(os.Stdout)
	} else {
		output = bufio.NewWriter(outStream)
	}

	fmt.Printf("%v => [task %v] => %v\n", inStream, taskNo, outStream)

	size, err := (*task).Execute(input, output)

	fmt.Println(size)

	if err != nil {
		log.Fatal(err)
		panic("execution error")
	}
}

func lookupTask(number int) (*AocTask, error) {
	switch number {
	case 1:
		return &Task1{}, nil

	default:
		return nil, &NoSuchTask{num: number}
	}
}
