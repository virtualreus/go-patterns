package mediator

import "fmt"

// паттерн проектирования, который позволяет уменьшить связанность множества классов между собой, благодаря перемещению этих связей в один класс-посредник.

// Посредник
type ControlTower struct {
	planes []*Plane
}

func (c *ControlTower) RequestTakeoff(plane *Plane) {
	fmt.Printf("Диспетчер: %s запрашивает взлет\n", plane.Name)

	// Проверяем, можно ли взлететь
	canTakeoff := true
	for _, p := range c.planes {
		if p.IsFlying {
			canTakeoff = false
			break
		}
	}

	if canTakeoff {
		fmt.Printf("Диспетчер: %s, взлет разрешен\n", plane.Name)
		plane.Takeoff()
	} else {
		fmt.Printf("Диспетчер: %s, ждите, взлетная полоса занята\n", plane.Name)
	}
}

func (c *ControlTower) RequestLanding(plane *Plane) {
	fmt.Printf("Диспетчер: %s запрашивает посадку\n", plane.Name)

	// Проверяем, можно ли приземлиться
	canLand := true
	for _, p := range c.planes {
		if p.IsLanding {
			canLand = false
			break
		}
	}

	if canLand {
		fmt.Printf("Диспетчер: %s, посадка разрешена\n", plane.Name)
		plane.Land()
	} else {
		fmt.Printf("Диспетчер: %s, ждите, полоса занята\n", plane.Name)
	}
}

// Самолет
type Plane struct {
	Name      string
	IsFlying  bool
	IsLanding bool
	Tower     *ControlTower
}

func (p *Plane) RequestTakeoff() {
	p.Tower.RequestTakeoff(p)
}

func (p *Plane) RequestLanding() {
	p.Tower.RequestLanding(p)
}

func (p *Plane) Takeoff() {
	p.IsFlying = true
	fmt.Printf("%s взлетел\n", p.Name)
}

func (p *Plane) Land() {
	p.IsFlying = false
	fmt.Printf("%s приземлился\n", p.Name)
}

// Использование
func main() {
	tower := &ControlTower{}

	planes := []*Plane{
		{Name: "Боинг 747", Tower: tower},
		{Name: "Аэробус A320", Tower: tower},
		{Name: "Ту-154", Tower: tower},
	}

	tower.planes = planes

	// Самолеты запрашивают взлет
	planes[0].RequestTakeoff() // Разрешен
	planes[1].RequestTakeoff() // Отказано (полоса занята)

	// Первый самолет садится
	planes[0].RequestLanding()

	// Теперь второй может взлететь
	planes[1].RequestTakeoff()
}
