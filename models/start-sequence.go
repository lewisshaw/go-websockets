package models

import "time"

type StartSequence interface {
	HootChannel() chan int
	CompleteChannel() chan bool
	TickChannel() chan int
	Start()
	IsComplete() bool
	GetRemainingTime() int
}

type startSequence struct {
	remainingTime   int
	hooterTimes     []int
	hootChannel     chan int
	completeChannel chan bool
	tickChannel     chan int
}

func NewStartSequence(countdownTime int, hooterTime []int) *startSequence {
	return &startSequence{
		remainingTime:   countdownTime,
		hooterTimes:     hooterTime,
		hootChannel:     make(chan int),
		completeChannel: make(chan bool),
		tickChannel:     make(chan int),
	}
}

func (ss *startSequence) HootChannel() chan int {
	return ss.hootChannel
}

func (ss *startSequence) CompleteChannel() chan bool {
	return ss.completeChannel
}

func (ss *startSequence) TickChannel() chan int {
	return ss.tickChannel
}

func (ss *startSequence) IsComplete() bool {
	return ss.remainingTime == 0
}

func (ss *startSequence) GetRemainingTime() int {
	return ss.remainingTime
}

func (ss *startSequence) Start() {
	ticker := time.NewTicker(time.Second)
	go ss.tick(ticker)
}

func (ss *startSequence) tick(t *time.Ticker) {
	ss.remainingTime--
	for ss.remainingTime >= 0 {
		<-t.C
		ss.tickChannel <- ss.remainingTime
		for _, hootTime := range ss.hooterTimes {
			if hootTime == ss.remainingTime {
				ss.hootChannel <- ss.remainingTime
			}
		}
		ss.remainingTime--
	}
	ss.completeChannel <- true
}
