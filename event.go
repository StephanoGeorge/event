package event

import (
	"sync"
	"sync/atomic"
)

type event struct {
	isSet     int32
	waitCount int32
	waitGroup sync.WaitGroup
}

func New() *event {
	return &event{}
}

// Set makes Wait() block
func (e *event) Set() {
	if atomic.LoadInt32(&e.isSet) == 0 && atomic.LoadInt32(&e.waitCount) == 0 {
		atomic.StoreInt32(&e.isSet, 1)
		e.waitGroup.Add(1)
	}
}

func (e *event) IsSet() bool {
	// Use Mutex to make read operation memory synchronization
	return atomic.LoadInt32(&e.isSet) != 0
}

// Clear makes Wait() not block
func (e *event) Clear() {
	if atomic.LoadInt32(&e.isSet) != 0 {
		atomic.StoreInt32(&e.isSet, 0)
		e.waitGroup.Done()
	}
}

// Wait blocks until Clear() called
func (e *event) Wait() {
	atomic.AddInt32(&e.waitCount, 1)
	e.waitGroup.Wait()
	atomic.AddInt32(&e.waitCount, -1)
}
