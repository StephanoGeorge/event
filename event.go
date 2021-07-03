package event

import (
	"sync"
)

type Event struct {
	waitGroup sync.WaitGroup
	isSet     bool
	waitCount sync.WaitGroup
	WaitChan  chan struct{}
	lock      sync.RWMutex
}

func New(withChan ...bool) *Event {
	e := &Event{}
	if len(withChan) > 0 && withChan[0] {
		e.WaitChan = make(chan struct{})
		go e.handleWaitChan()
	}
	return e
}

// Set makes Wait() block
func (e *Event) Set() {
	e.lock.Lock()
	if !e.isSet {
		e.waitCount.Wait()
		e.isSet = true
		e.waitGroup.Add(1)
	}
	e.lock.Unlock()
}

func (e *Event) IsSet() bool {
	e.lock.RLock()
	isSet := e.isSet
	e.lock.RUnlock()
	return isSet
}

// Clear makes Wait() not block
func (e *Event) Clear() {
	e.lock.Lock()
	if e.isSet {
		e.isSet = false
		e.waitGroup.Done()
	}
	e.lock.Unlock()
}

// Wait blocks until Clear() called
func (e *Event) Wait() {
	e.waitCount.Add(1)
	e.waitGroup.Wait()
	e.waitCount.Done()
}

func (e *Event) handleWaitChan() {
	for {
		e.Wait()
		e.WaitChan <- struct{}{}
	}
}
