package event

import (
	"sync"
)

type Event struct {
	WaitChan     chan struct{}
	waitGroup    sync.WaitGroup
	waitCount    sync.WaitGroup
	isSet        bool
	isClosed     bool
	lock         sync.RWMutex
	isClosedLock sync.RWMutex
}

func New(withChan ...bool) *Event {
	e := &Event{}
	e.Set()
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
		// wait for all Event waiter to finish, because number of waitGroup waiter must be 0 when Adding
		e.waitCount.Wait()
		e.isSet = true
		e.waitGroup.Add(1)
	}
	e.lock.Unlock()
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
	// avoid racing with Set()
	e.lock.RLock()
	e.waitCount.Add(1)
	e.lock.RUnlock()
	e.waitGroup.Wait()
	e.waitCount.Done()
}

func (e *Event) IsSet() bool {
	e.lock.RLock()
	isSet := e.isSet
	e.lock.RUnlock()
	return isSet
}

func (e *Event) Close() {
	e.Clear()
	e.isClosedLock.Lock()
	e.isClosed = true
	e.isClosedLock.Unlock()
	if e.WaitChan != nil {
		<-e.WaitChan
	}
}

func (e *Event) handleWaitChan() {
	for {
		e.Wait()
		// a struct must have be sent before breaking
		e.WaitChan <- struct{}{}
		e.isClosedLock.RLock()
		isClosed := e.isClosed
		e.isClosedLock.RUnlock()
		if isClosed {
			break
		}
	}
}
