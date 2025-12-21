package main

import (
	"fmt"
	"sync"
)

// Паттерн который гарантирует что у типа есть только один инстанс и предоставляет к нему глобальную точку доступа.
// Можно заюзать как логгер, подключение к бд/пулу соединений, конфиг кэши, менеджеры доступа.

var mu = &sync.Mutex{}

type single struct{}

var singleInstance *single

func getInstanceDefault() *single {
	if singleInstance == nil {
		mu.Lock()
		defer mu.Unlock()
		if singleInstance == nil {
			fmt.Println("Creating single instance now.")
			singleInstance = &single{}
		} else {
			fmt.Println("Single instance already created.")
		}
	} else {
		fmt.Println("Single instance already created.")
	}

	return singleInstance
}

// Пример с sync.Once

var once sync.Once

func getInstanceOnce() *single {
	once.Do(func() {
		fmt.Println("Creating single instance now.")
		singleInstance = &single{}
	})
	return singleInstance
}

func main() {

	for i := 0; i < 30; i++ {
		go getInstanceOnce()
	}

	// Scanln is similar to Scan, but stops scanning at a newline and
	// after the final item there must be a newline or EOF.
	fmt.Scanln()
}
