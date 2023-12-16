package main

import (
	"fmt"
	"hill-climbing/runner"
	"os"
	"os/signal"
	"syscall"
)

func main() {

	file, err := os.Open("data/dataset6-sorted.txt")
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
		IterationCount:  70000,
		Data:            &data,
		PerfectData:     &perfect,
		CurrentSequence: ss,
		BestSequence:    ss,
	}

	registerSignalHandler(&r.BestSequence)
	r.Run()
}

func registerSignalHandler(state **runner.State) {
	signal_chan := make(chan os.Signal, 1)
	signal.Notify(signal_chan, syscall.SIGINT)

	go func() {
		for {
			s := <-signal_chan
			if s == syscall.SIGINT {
				fmt.Printf("%#v", **state)
				os.Exit(0)
			}
		}
	}()
}
