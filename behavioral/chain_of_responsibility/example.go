package chain_of_responsibility

import "fmt"

/*
поведенческий паттерн проектирования, который позволяет передавать
запросы последовательно по цепочке обработчиков.
Каждый последующий обработчик решает, может ли он обработать
запрос сам и стоит ли передавать запрос дальше по цепи.
*/

// Интерфейс обработчика
type Handler interface {
	Handle(request string) string
	SetNext(handler Handler)
}

// Базовый обработчик
type BaseHandler struct {
	next Handler
}

func (b *BaseHandler) SetNext(handler Handler) {
	b.next = handler
}

// Конкретные обработчики
type AuthHandler struct {
	BaseHandler
}

func (a *AuthHandler) Handle(request string) string {
	if request == "auth" {
		return "Аутентификация выполнена"
	}

	if a.next != nil {
		return a.next.Handle(request)
	}

	return "Не обработано"
}

type LoggingHandler struct {
	BaseHandler
}

func (l *LoggingHandler) Handle(request string) string {
	if request == "log" {
		return "Логирование выполнено"
	}

	if l.next != nil {
		return l.next.Handle(request)
	}

	return "Не обработано"
}

type CacheHandler struct {
	BaseHandler
}

func (c *CacheHandler) Handle(request string) string {
	if request == "cache" {
		return "Кэширование выполнено"
	}

	if c.next != nil {
		return c.next.Handle(request)
	}

	return "Не обработано"
}

// Использование
func main() {
	// Создаем цепочку
	auth := &AuthHandler{}
	logging := &LoggingHandler{}
	cache := &CacheHandler{}

	// Настраиваем порядок
	auth.SetNext(logging)
	logging.SetNext(cache)

	// Запросы идут по цепочке
	fmt.Println(auth.Handle("auth"))    // "Аутентификация выполнена"
	fmt.Println(auth.Handle("log"))     // "Логирование выполнено"
	fmt.Println(auth.Handle("cache"))   // "Кэширование выполнено"
	fmt.Println(auth.Handle("unknown")) // "Не обработано"
}
