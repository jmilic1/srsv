package main

import (
	"context"
	"os"
	"os/signal"
	"strconv"
	"syscall"
	"time"

	"lab3/components"
	"lab3/parameters"
)

func main() {
	inputNum := parameters.Inputs

	inputs := make([]*components.Input, inputNum)
	for i := 0; i < inputNum; i++ {
		answerChan := make(chan int)
		inputs[i] = components.NewInput(i, parameters.InputPeriod, parameters.InputPeriod*2, answerChan)
	}

	ctx, cancel := context.WithCancel(context.Background())
	for _, input := range inputs {
		go components.Driver(input, ctx)
		go input.Start(ctx)
	}

	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, syscall.SIGINT)

	_ = <-signalChan
	cancel()

	numChanges := 0
	var sumAnswerDuration time.Duration
	var maxAnswerDuration time.Duration
	problems := 0

	for _, input := range inputs {
		println("thread " + strconv.Itoa(input.Id) + ": number of state changes: " + strconv.Itoa(input.NumberStateChanges) + "\n\t" +
			"sum of answer durations: " + input.SumAnswerDuration.String() + "\n\t" +
			"max answer duration: " + input.MaxAnswerDuration.String() + "\n\t" +
			"number of problems: " + strconv.Itoa(input.NumberProblems))

		numChanges += input.NumberStateChanges
		sumAnswerDuration += input.SumAnswerDuration
		problems += input.NumberProblems

		if maxAnswerDuration < input.MaxAnswerDuration {
			maxAnswerDuration = input.MaxAnswerDuration
		}
	}

	println("whole system: number of state changes: " + strconv.Itoa(numChanges) + "\n\t" +
		"sum of answer durations: " + sumAnswerDuration.String() + "\n\t" +
		"max answer duration: " + maxAnswerDuration.String() + "\n\t" +
		"number of problems: " + strconv.Itoa(problems))
}
