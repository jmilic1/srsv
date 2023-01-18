package components

import (
	"context"
	"math/rand"
	"strconv"
	"time"
)

// Driver should be run before simulation threads
func Driver(input *Input, ctx context.Context) {
	lastStates := make(map[int]int, 0)

	for {
		select {
		case <-ctx.Done():
			return
		default:
			state := input.state
			lastState, exists := lastStates[input.Id]

			if !exists {
				lastState = -1
			}

			if state == lastState {
				println("upr: input " + strconv.Itoa(input.Id) + ": no new change (" + strconv.Itoa(state) + "). Sleeping")
				time.Sleep(time.Millisecond * 30)
				continue
			}

			println("upr: input " + strconv.Itoa(input.Id) + ": change (" + strconv.Itoa(lastState) + "->" + strconv.Itoa(state) + "), processing")
			simulateProcess()
			input.AnswerChan <- rand.Intn(100)
			lastStates[input.Id] = state
			println("upr: input " + strconv.Itoa(input.Id) + ": end of process, set to (" + strconv.Itoa(state) + ")")
		}
	}
}

func simulateProcess() {
	random := rand.Intn(100)

	var prod int64

	switch true {
	case random < 50:
		for i := 0; i < numIter*0.1; i++ {
			prod *= int64(i)
		}
	case random < 80:
		for i := 0; i < numIter*0.2; i++ {
			prod *= int64(i)
		}
	case random < 95:
		for i := 0; i < numIter*0.4; i++ {
			prod *= int64(i)
		}
	case random < 100:
		for i := 0; i < numIter*0.7; i++ {
			prod *= int64(i)
		}
	}
}
