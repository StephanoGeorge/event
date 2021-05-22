package event

import (
	"testing"
	"time"
)

func TestEvent(t *testing.T) {
	e := &Event{}
	if e.IsSet() {
		panic("Fail")
	}

	for i := 0; i < 3; i++ {
		e.Set()
		if !e.IsSet() {
			panic("Fail")
		}
		c := make(chan struct{})
		go func() {
			e.Wait()
			c <- struct{}{}
		}()
		select {
		case <-c:
			panic("Fail")
		case <-time.After(time.Second):
		}
		e.Clear()
		if e.IsSet() {
			panic("Fail")
		}
		select {
		case <-c:
		case <-time.After(time.Second):
			panic("Fail")
		}
	}

	go func() {
		for i := 0; i < 100; i++ {
			e.Set()
		}
	}()
	go func() {
		for i := 0; i < 100; i++ {
			e.Clear()
		}
	}()
	go func() {
		for i := 0; i < 100; i++ {
			e.Wait()
		}
	}()
}
