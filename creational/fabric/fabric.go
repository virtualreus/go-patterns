// Порождающий паттерн проектирования, который определяет общий интерфейс для обьектов, позволяя подтипам изменять тип создаваемых обьектов
// Нужен когда мы не знаем какие конкретно инстансы типов нам надо создавать, а так же в гибкости и сокрытие сложности создания
package main

import "fmt"

// IGun Интерфейс общий для методов оружия
type IGun interface {
	setName(name string)
	setPower(power int)
	getName() string
	getPower() int
}

// Gun Конкретный продукт
type Gun struct {
	name  string
	power int
}

func (g *Gun) setName(name string) {
	g.name = name
}

func (g *Gun) getName() string {
	return g.name
}

func (g *Gun) setPower(power int) {
	g.power = power
}

func (g *Gun) getPower() int {
	return g.power
}

// Ak47 - конкретный продукт
type Ak47 struct {
	Gun
}

func newAk47() IGun {
	return &Ak47{
		Gun: Gun{
			name:  "AK47 gun",
			power: 4,
		},
	}
}

// musket конкретный продукт
type musket struct {
	Gun
}

func newMusket() IGun {
	return &musket{
		Gun: Gun{
			name:  "Musket gun",
			power: 1,
		},
	}
}

// getGun - сама фабрика
func getGun(gunType string) (IGun, error) {
	if gunType == "ak47" {
		return newAk47(), nil
	}
	if gunType == "musket" {
		return newMusket(), nil
	}
	return nil, fmt.Errorf("wrong gun")
}

func main() {
	ak47Gun, _ := getGun("ak47")
	musketGun, _ := getGun("musket")

	printDetails(ak47Gun)
	printDetails(musketGun)
}

func printDetails(g IGun) {
	fmt.Printf("Gun: %s", g.getName())
	fmt.Println()
	fmt.Printf("Power: %d", g.getPower())
	fmt.Println()
}
