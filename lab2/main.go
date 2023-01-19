package main

import (
	"context"
	"os"
	"os/signal"
	"strconv"
	"syscall"
	"time"

	"lab2/components"
)

func main() {
	inputs := components.NewInputsStruct()

	answer := &components.Answer{
		SecondPeriod: 0,
		Stopped:      0,
	}
	ctx, cancel := context.WithCancel(context.Background())
	go components.Driver(inputs, nil, ctx, answer)

	for _, input := range inputs.InputArr {
		go input.Start(ctx)
	}

	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, syscall.SIGINT)

	for {
		b := false
		select {
		case <-signalChan:
			b = true
		default:
			time.Sleep(time.Millisecond * 100)
			components.ProcessTask(inputs, answer)
		}
		if b {
			break
		}
	}

	println("cancelling")
	cancel()

	numChanges := 0
	var sumAnswerDuration time.Duration
	var maxAnswerDuration time.Duration
	problems := 0

	for _, input := range inputs.InputArr {
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

	println("upr stats: second period", answer.SecondPeriod, "\n\t"+
		"stopped: ", answer.Stopped, "\n\t")
}
