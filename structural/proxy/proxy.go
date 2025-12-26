package proxy

import (
	"fmt"
	"sync"
	"time"
)

// Прокси - паттерн, который предоставляет обьект заместитель который контроллирует доступ к основному обьекту. Как обертка вокруг реального обьекта

type Subject interface {
	Request() string
}

// Реальный объект (создается дорого)
type RealSubject struct{}

func NewRealSubject() *RealSubject {
	// Имитация дорогой инициализации
	time.Sleep(2 * time.Second)
	fmt.Println("RealSubject created")
	return &RealSubject{}
}

func (r *RealSubject) Request() string {
	return "RealSubject: Handling request"
}

type Proxy struct {
	realSubject *RealSubject
	mu          sync.Mutex
}

func (p *Proxy) Request() string {
	p.mu.Lock()
	defer p.mu.Unlock()

	if p.realSubject == nil {
		p.realSubject = NewRealSubject()
	}

	return "Proxy: " + p.realSubject.Request()
}
