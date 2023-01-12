package components

import (
	"context"
	"math/rand"
	"strconv"
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
				println("upr: input " + strconv.Itoa(input.Id) + ": no new change (" + strconv.Itoa(state) + ")")
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
	var prod int64
	for i := 0; i < numIter; i++ {
		prod *= int64(i)
	}

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
}
