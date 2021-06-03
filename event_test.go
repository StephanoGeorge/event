package event

import (
	"math/rand"
	"sync"
	"testing"
	"time"
)

func TestEvent(t *testing.T) {
	e := New()
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

	count := 30
	w := sync.WaitGroup{}
	w.Add(count)
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < count; i++ {
		go func() {
			time.Sleep(time.Duration(rand.Intn(1000000)))
			e.Set()
		}()
	}
	for i := 0; i < count; i++ {
		go func() {
			time.Sleep(time.Duration(rand.Intn(1000000)))
			e.Clear()
		}()
	}
	for i := 0; i < count; i++ {
		go func() {
			time.Sleep(time.Duration(rand.Intn(1000000)))
			e.Wait()
			w.Done()
		}()
	}
	w.Wait()
	if e.IsSet() {
		panic("Fail")
	}
	if e.waitCount != 0 {
		panic("Fail")
	}
}
