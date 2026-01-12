package main

import (
	"fmt"
	"sync"
)

func SplitChannels[T any](inputCh <-chan T, n int) []<-chan T {
	outputChs := make([]chan T, n)
	for i := 0; i < n; i++ {
		outputChs[i] = make(chan T)
	}

	go func() {
		for value := range inputCh {
			for i := 0; i < n; i++ {
				outputChs[i] <- value
			}
		}
		for _, ch := range outputChs {
			close(ch)
		}
	}()

	resultChs := make([]<-chan T, n)
	for i := 0; i < n; i++ {
		resultChs[i] = outputChs[i]
	}
	return resultChs
}

func main() {
	channel := make(chan int)
	go func() {
		defer close(channel)
		for i := 0; i < 10; i++ {
			channel <- i
		}
	}()

	channels := SplitChannels(channel, 2)

	var wg sync.WaitGroup
	wg.Add(2)

	go func() {
		defer wg.Done()
		for value := range channels[0] {
			fmt.Println("ch1: ", value)
		}
	}()
	go func() {
		defer wg.Done()
		for value := range channels[1] {
			fmt.Println("ch2: ", value)
		}
	}()

	wg.Wait()
}
