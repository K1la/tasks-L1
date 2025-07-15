package main

import (
	"fmt"
	"math"
)

/*
	Расстояние между точками

	Разработать программу нахождения расстояния между
	двумя точками на плоскости. Точки представлены
	в виде структуры Point с инкапсулированными
	(приватными) полями x, y (типа float64) и конструктором.
	Расстояние рассчитывается по формуле между координатами двух точек.
*/

type Point struct {
	x, y float64
}

func NewPoint(x, y float64) *Point {
	return &Point{x, y}
}

func (p *Point) Distance(other *Point) float64 {
	dx := other.x - p.x
	dy := other.y - p.y
	return math.Hypot(dx, dy)
}

func main() {
	p1 := NewPoint(1, 2)
	p2 := NewPoint(3, 4)

	d := p1.Distance(p2)

	fmt.Printf("p1: (%.2f, %.2f)\n", p1.x, p1.y)
	fmt.Printf("p2: (%.2f, %.2f)\n", p2.x, p2.y)
	fmt.Printf("Расстояние между p1 и p2 = %.8f\n", d)
}
