package components

import (
	"context"
	"math/rand"
	"time"
)

var lastStates map[int]int

var taskInProcess int
var jobInProcess int64
var period Period
var noOver int

type Answer struct {
	SecondPeriod int
	Stopped      int
}

// Driver should be run before simulation threads
func Driver(inputs *InputsStruct, requestChans []<-chan int, ctx context.Context, answer *Answer) {

	taskInProcess = -1 // zadatak_u_obradi
	jobInProcess = 0   // posao_u_obradi
	period = Period(0)

	noOver = 0

	lastStates = make(map[int]int, 0)

	for {
		select {
		case <-ctx.Done():
			return
		default:
			ProcessTask(inputs, answer)
			// for _, input := range inputs.InputArr {
			// 	state := input.state
			// 	lastState, exists := lastStates[input.Id]
			//
			// 	if !exists {
			// 		lastState = -1
			// 	}
			//
			// 	if state == lastState {
			// 		println("upr: input " + strconv.Itoa(input.Id) + ": no new change (" + strconv.Itoa(state) + ")")
			// 		continue
			// 	}
			//
			// 	println("upr: input " + strconv.Itoa(input.Id) + ": change (" + strconv.Itoa(lastState) + "->" + strconv.Itoa(state) + "), processing")
			// 	simulateProcess()
			// 	input.AnswerChan <- rand.Intn(100)
			// 	lastStates[input.Id] = state
			// 	println("upr: input " + strconv.Itoa(input.Id) + ": end of process, set to (" + strconv.Itoa(state) + ")")
			// }
		}
	}
}

func ProcessTask(inputs *InputsStruct, answer *Answer) {
	if taskInProcess != -1 {
		if period == FirstPeriod && noOver >= 10 {
			answer.SecondPeriod++
			println("upr: accepting second period to task ", taskInProcess)

			period = SecondPeriod
			noOver = 0
			return
		} else {
			answer.Stopped++
			println("upr: stopping task ", taskInProcess)
		}
	}

	takeNext := 1
	var myTask int
	var myJob int64
	for takeNext == 1 {
		takeNext = 0

		taskInProcess = getNext(inputs)
		myTask = taskInProcess

		lastState, exists := lastStates[myTask]

		if myTask == -1 || (inputs.InputArr[myTask].state == lastState && exists) {
			// nema promjene ulaza za ovaj zadatak
			taskInProcess = -1
			jobInProcess = 0
			return
		}

		jobInProcess = time.Now().UnixMilli()
		myJob = jobInProcess
		period = FirstPeriod

		println("upr: begin processing task ", myTask)

		processingTime := simulateProcess()
		for myJob == jobInProcess && processingTime > 0 {
			time.Sleep(time.Millisecond * 5)
			processingTime = processingTime - time.Millisecond*5
		}

		// obrada je ili završena ili je prekinuta
		if myJob == jobInProcess && processingTime <= 0 {
			// zabilježi kraj obrade
			inputs.InputArr[myTask].AnswerChan <- inputs.InputArr[myTask].state
			println("upr: process finished for task ", myTask)

			taskInProcess = -1
			jobInProcess = 0
			if period == FirstPeriod {
				noOver++
			} else {
				takeNext = 1
			}
		}
	}
}

func getNext(inputs *InputsStruct) int {
	nextTask := 0
	tip := inputs.Queue[inputs.T]
	inputs.T = (inputs.T + 1) % 10
	if tip != -1 {
		nextTask = inputs.InputNums[tip][inputs.Ind[tip]]
		inputs.Ind[tip] = (inputs.Ind[tip] + 1) % inputs.Maxes[tip]
	}
	return nextTask - 1
}

func simulateProcess() time.Duration {
	// random := rand.Intn(100)
	// switch true {
	// case random < 20:
	// 	<-time.After(time.Millisecond * 30)
	// case random >= 20 && random < 70:
	// 	<-time.After(time.Millisecond * 50)
	// case random >= 70 && random < 95:
	// 	<-time.After(time.Millisecond * 80)
	// case random > 95:
	// 	<-time.After(time.Millisecond * 120)
	// }
	random := rand.Intn(100)
	switch true {
	case random < 20:
		return time.Millisecond * 30
	case random >= 20 && random < 70:
		return time.Millisecond * 50
	case random >= 70 && random < 95:
		return time.Millisecond * 80
	case random > 95:
		return time.Millisecond * 120
	}

	return 0
}
