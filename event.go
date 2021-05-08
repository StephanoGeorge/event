package event

import "sync"

type Event struct {
	isSet     bool
	waitGroup sync.WaitGroup
	lock      sync.RWMutex
}

// Set makes Wait() block
func (e *Event) Set() {
	e.lock.Lock()
	defer e.lock.Unlock()
	if !e.isSet {
		e.isSet = true
		e.waitGroup.Add(1)
	}
}

func (e *Event) IsSet() bool {
	// Use Mutex to make read operation memory synchronization
	e.lock.RLock()
	defer e.lock.RUnlock()
	return e.isSet
}

// Clear makes Wait() not block
func (e *Event) Clear() {
	e.lock.Lock()
	defer e.lock.Unlock()
	if e.isSet {
		e.isSet = false
		e.waitGroup.Done()
	}
}

// Wait blocks until Clear() called
func (e *Event) Wait() {
	e.waitGroup.Wait()
}
