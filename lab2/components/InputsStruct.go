package components

import "lab2/parameters"

type InputsStruct struct {
	InputArr     []*Input
	RequestChans []<-chan int

	InputNums [][]int
	Queue     []int
	T         int
	Ind       []int
	Maxes     []int
}

func NewInputsStruct() *InputsStruct {
	length := 3 + (18 - 3) + (38 - 18 - 3)
	inputs := make([]*Input, length)
	requestChans := make([]<-chan int, length)
	for i := 0; i < length; i++ {
		answerChan := make(chan int)
		requestChan := make(chan int)
		requestChans = append(requestChans, requestChan)
		inputs[i] = NewInput(i, parameters.InputPeriod, parameters.InputPeriod*2, answerChan, requestChan)
	}

	inputNums := make([][]int, 3)
	first := []int{1, 2, 3}
	second := []int{4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18}
	third := []int{19, 20, 21, 22, 23, 24, 25, 26, 27, 28, 29, 30, 31, 32, 33, 34, 35, 36, 37, 38}

	inputNums[0] = first
	inputNums[1] = second
	inputNums[2] = third

	queue := []int{0, 1, -1, 0, 1, -1, 0, 1, 2, -1}
	k := 0
	ind := []int{0, 0, 0}
	maxes := []int{3, 15, 20}

	return &InputsStruct{
		InputArr:     inputs,
		RequestChans: requestChans,
		InputNums:    inputNums,
		Queue:        queue,
		T:            k,
		Ind:          ind,
		Maxes:        maxes,
	}
}
