package main

import (
	"fmt"
	"hill-climbing/runner"
	"os"
)

func main() {

	file, err := os.Open("data/dataset6.txt")
	if err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
		return
	}
	data, err := runner.ReadInts(file)
	if err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
		return
	}

	perfect := runner.CalculatePerfect(&data)
	ss := runner.InitialState(&data, &perfect)

	r := runner.Runner{
		IterationCount:  10000000,
		Data:            &data,
		PerfectData:     &perfect,
		CurrentSequence: ss,
		BestSequence:    ss,
	}

	r.Run()
}
