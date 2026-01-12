package main

import (
	"fmt"
	"sync"
)

func MergeChannels[T any](chans ...chan T) <-chan T {
	var wg sync.WaitGroup
	wg.Add(len(chans))

	outputCh := make(chan T)
	for _, channel := range chans {
		go func() {
			defer wg.Done()
			for v := range channel {
				outputCh <- v
			}
		}()
	}

	go func() {
		wg.Wait()
		close(outputCh)
	}()
	return outputCh
}

// Смержить n каналов в 1 канал
func main() {
	ch1 := make(chan int)
	ch2 := make(chan int)
	ch3 := make(chan int)

	go func() {
		defer func() {
			close(ch1)
			close(ch2)
			close(ch3)
		}()

		for i := 0; i < 100; i += 3 {
			ch1 <- i
			ch1 <- i + 1
			ch1 <- i + 2
		}
	}()

	for v := range MergeChannels(ch1, ch2, ch3) {
		fmt.Println(v)
	}
}
