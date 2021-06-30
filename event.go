package event

import (
	"sync"
	"sync/atomic"
)

type Event struct {
	isSet     int32
	waitCount int32
	waitGroup sync.WaitGroup
}

func New() *Event {
	return &Event{}
}

// Set makes Wait() block
func (e *Event) Set() {
	if atomic.LoadInt32(&e.isSet) == 0 && atomic.LoadInt32(&e.waitCount) == 0 {
		atomic.StoreInt32(&e.isSet, 1)
		e.waitGroup.Add(1)
	}
}

func (e *Event) IsSet() bool {
	return atomic.LoadInt32(&e.isSet) != 0
}

// Clear makes Wait() not block
func (e *Event) Clear() {
	if atomic.LoadInt32(&e.isSet) != 0 {
		atomic.StoreInt32(&e.isSet, 0)
		e.waitGroup.Done()
	}
}

// Wait blocks until Clear() called
func (e *Event) Wait() {
	atomic.AddInt32(&e.waitCount, 1)
	e.waitGroup.Wait()
	atomic.AddInt32(&e.waitCount, -1)
}
