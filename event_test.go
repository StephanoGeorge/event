package event

import (
	"sync"
	"testing"
	"time"
)

func TestEvent(t *testing.T) {
	e := &Event{}
	if e.IsSet() {
		panic("Fail")
	}

	for i := 0; i < 3; i++ {
		w := &sync.WaitGroup{}
		for i := 0; i < 3; i++ {
			w.Add(1)
			go func() {
				e.Set()
				w.Done()
			}()
		}

		w.Wait()
		if !e.IsSet() {
			panic("Fail")
		}

		c := make(chan struct{})
		for i := 0; i < 3; i++ {
			go func() {
				e.Wait()
				c <- struct{}{}
			}()
		}
		time.Sleep(time.Second)
		if len(c) != 0 {
			panic("Fail")
		}
		for i := 0; i < 3; i++ {
			w.Add(1)
			go func() {
				e.Clear()
				w.Done()
			}()
		}
		w.Wait()
		if e.IsSet() {
			panic("Fail")
		}
		for i := 0; i < 3; i++ {
			select {
			case <-time.After(time.Second):
				panic("Fail")
			case <-c:
			}
		}
	}
}
