package command

import "fmt"

/*
поведенческий паттерн проектирования, который превращает запросы в объекты,
позволяя передавать их как аргументы при вызове методов,
ставить запросы в очередь, логировать их, а также поддерживать отмену операций.
*/

// Команда (интерфейс)
type Command interface {
	Execute()
}

// Получатель (тот, кто выполняет действие)
type Light struct{}

func (l *Light) TurnOn() {
	fmt.Println("Свет включен")
}

func (l *Light) TurnOff() {
	fmt.Println("Свет выключен")
}

// Конкретная команда "Включить свет"
type TurnOnCommand struct {
	light *Light
}

func (c *TurnOnCommand) Execute() {
	c.light.TurnOn()
}

// Конкретная команда "Выключить свет"
type TurnOffCommand struct {
	light *Light
}

func (c *TurnOffCommand) Execute() {
	c.light.TurnOff()
}

// Инициатор (тот, кто вызывает команду)
type Switch struct {
	command Command
}

func (s *Switch) Press() {
	s.command.Execute()
}

// Использование
func main() {
	// Создаем свет
	light := &Light{}

	// Создаем команды
	turnOn := &TurnOnCommand{light: light}
	turnOff := &TurnOffCommand{light: light}

	// Создаем выключатель
	sw := &Switch{}

	// Включаем свет
	sw.command = turnOn
	sw.Press() // "Свет включен"

	// Выключаем свет
	sw.command = turnOff
	sw.Press() // "Свет выключен"
}
