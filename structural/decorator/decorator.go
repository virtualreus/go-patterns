package decorator

// Паттерн декоратор - структурный паттерн который
// позволяет динамически добавлять обьектам новую функциональность
// оборачивая их в полезные обертки

// нужен для модификации в рантайме

// Pizza Базовый интерфейс
type Pizza interface {
	GetDescription() string
	GetCost() float64
}

// Margherita - конкретная реализация
type Margherita struct{}

func (m *Margherita) GetDescription() string {
	return "Пицца Маргарита"
}

func (m *Margherita) GetCost() float64 {
	return 5.0
}

type PizzaDecorator struct {
	pizza Pizza
}

func (p *PizzaDecorator) GetDescription() string {
	return p.pizza.GetDescription()
}

func (p *PizzaDecorator) GetCost() float64 {
	return p.pizza.GetCost()
}

// Конкретные декораторы
type PepperoniDecorator struct {
	PizzaDecorator
}

func NewPepperoniDecorator(pizza Pizza) *PepperoniDecorator {
	return &PepperoniDecorator{PizzaDecorator{pizza: pizza}}
}

func (p *PepperoniDecorator) GetDescription() string {
	return p.pizza.GetDescription() + ", пепперони"
}

func (p *PepperoniDecorator) GetCost() float64 {
	return p.pizza.GetCost() + 2.0
}
