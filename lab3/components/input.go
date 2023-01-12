package components

import (
	"context"
	"math/rand"
	"strconv"
	"time"

	"lab3/parameters"
)

const minContinous = 10

type Period int

const (
	NotRunning Period = iota
	FirstPeriod
	SecondPeriod
)

const numIter = 10e6

type Input struct {
	Id         int
	period     time.Duration
	startAfter time.Duration

	state               int
	lastChangeTimestamp time.Time
	lastAnswer          int
	lastAnswerTimestamp time.Time

	AnswerChan  chan int
	RequestChan chan<- int

	NumberStateChanges int
	SumAnswerDuration  time.Duration
	NumberProblems     int
	MaxAnswerDuration  time.Duration
}

func NewInput(id int, period time.Duration, startAfter time.Duration, answerChan chan int, requestChan chan<- int) *Input {
	return &Input{
		Id:          id,
		period:      period,
		startAfter:  startAfter,
		AnswerChan:  answerChan,
		RequestChan: requestChan,
	}
}

func (i *Input) Start(ctx context.Context) {
	<-time.After(i.startAfter)

	for {
		select {
		case <-ctx.Done():
			return
		default:
			i.do()
		}
	}
}

func (i *Input) do() {

	for {
		i.state = 100 + rand.Intn(900)
		i.lastChangeTimestamp = time.Now()
		i.NumberStateChanges++

		println("thread " + strconv.Itoa(i.Id) + ": state change " + strconv.Itoa(i.state))

		select {
		case answer := <-i.AnswerChan:
			i.lastAnswer = answer
			i.lastAnswerTimestamp = time.Now()
			diff := i.lastAnswerTimestamp.Sub(i.lastChangeTimestamp)
			if diff > i.period {
				i.NumberProblems++
			}
			println("thread " + strconv.Itoa(i.Id) + ": answered with " + strconv.Itoa(answer))
			i.SumAnswerDuration += diff
			if diff > i.MaxAnswerDuration {
				i.MaxAnswerDuration = diff
			}

			// postpone := i.period * time.Duration(rand.Float64()*(parameters.K-1))
			// time.Sleep(postpone + i.period - time.Now().Sub(i.lastChangeTimestamp))
			var prod int64
			for i := 0; i < numIter; i++ {
				prod *= int64(i)
			}

		default:
			println("thread " + strconv.Itoa(i.Id) + ": problem: not answered, sleeping")
			time.Sleep(parameters.InputSleep)
		}
	}
}
