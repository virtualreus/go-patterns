package main

import (
	"fmt"
	"sync"
	"time"
)

type Semaphore struct {
	tickets chan struct{}
}

func NewSemaphore(n int) Semaphore {
	return Semaphore{
		tickets: make(chan struct{}, n),
	}
}

func (s *Semaphore) Acquire() {
	s.tickets <- struct{}{}
}

func (s *Semaphore) Relase() {
	<-s.tickets
}

func main() {
	wg := sync.WaitGroup{}
	wg.Add(6)
	semaphore := NewSemaphore(5)
	for i := 0; i < 6; i++ {
		semaphore.Acquire()
		go func() {
			defer func() {
				wg.Done()
				semaphore.Relase()
			}()
			fmt.Println("working..")
			time.Sleep(time.Second * 2)
			fmt.Println("exiting..")
		}()
	}
	wg.Wait()
}
