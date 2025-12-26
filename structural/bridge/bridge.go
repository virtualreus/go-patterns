package bridge

import "fmt"

// Позволяет разделить один или несколько классов на две иерархии - абстракцию или реализацию

// Renderer Shape Базовые интерфейсы
type Renderer interface {
	RenderCircle(radius float64)
	RenderSquare(side float64)
}

type Shape interface {
	Draw()
	Resize(factor float64)
}

// Конкретные реализации
type VectorRenderer struct {
}

func (v *VectorRenderer) RenderCircle(radius float64) {
	fmt.Printf("Drawing circle with radius %.2f using vector graphics\n", radius)
}

func (v *VectorRenderer) RenderSquare(side float64) {
	fmt.Printf("Drawing square with side %.2f using vector graphics\n", side)
}

// Конкретные реализации
type RasterRenderer struct {
	DPI int
}

func (r *RasterRenderer) RenderCircle(radius float64) {
	fmt.Printf("Drawing circle with radius %.2f using raster graphics (DPI: %d)\n",
		radius, r.DPI)
}

func (r *RasterRenderer) RenderSquare(side float64) {
	fmt.Printf("Drawing square with side %.2f using raster graphics (DPI: %d)\n",
		side, r.DPI)
}

// Конкретные абстракции
type Circle struct {
	renderer Renderer
	radius   float64
}

func NewCircle(renderer Renderer, radius float64) *Circle {
	return &Circle{
		renderer: renderer,
		radius:   radius,
	}
}

func (c *Circle) Draw() {
	c.renderer.RenderCircle(c.radius)
}

func (c *Circle) Resize(factor float64) {
	c.radius *= factor
}
