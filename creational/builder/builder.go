package main

import "fmt"

// Паттерн билдер нужен, для того чтобы создавать сложные объекты пошагово. Отделяет конструирование одного сложного объекта от его представления
// Один и тот же процесс конструирования может создавать разные представления.
// Это базовая реализация, есть вариант с директором, но я эту ёбань даже разбирать не хочу

// Car структура автомобиля
type Car struct {
	Brand            string
	Color            string
	EnginePower      int
	MultimediaSystem string
}

// CarBuilder Builder - сам билдер автомобиля
type CarBuilder interface {
	SetBrand(brand string) CarBuilder
	SetColor(color string) CarBuilder
	SetEnginePower(power int) CarBuilder
	SetMultimediaSystem(system string) CarBuilder
	Build() Car
}

// Имплементация
type carBuilder struct {
	brand            string
	color            string
	enginePower      int
	multimediaSystem string
}

func NewCarBuilder() CarBuilder {
	return &carBuilder{}
}

// SetBrand Реализуем методы интерфейса, которые будут собирать наш объект
func (cb *carBuilder) SetBrand(brand string) CarBuilder {
	cb.brand = brand
	return cb
}

func (cb *carBuilder) SetColor(color string) CarBuilder {
	cb.color = color
	return cb
}

func (cb *carBuilder) SetEnginePower(power int) CarBuilder {
	cb.enginePower = power
	return cb
}

func (cb *carBuilder) SetMultimediaSystem(system string) CarBuilder {
	cb.multimediaSystem = system
	return cb
}
func (cb *carBuilder) Build() Car {
	return Car{
		Brand:            cb.brand,
		Color:            cb.color,
		EnginePower:      cb.enginePower,
		MultimediaSystem: cb.multimediaSystem,
	}
}

func main() {
	// Создаем нового строителя автомобиля
	builder := NewCarBuilder()

	// Добавляем различные характеристики
	car := builder.SetBrand("Toyota").
		SetColor("Черный").
		SetEnginePower(250).
		SetMultimediaSystem("Premium").
		Build()

	// Выводим созданный объект автомобиля
	fmt.Printf("Автомобиль бренда %s, цвета %s с мощностью двигателя %d и мультимедийной системой %s\n",
		car.Brand, car.Color, car.EnginePower, car.MultimediaSystem)
}
