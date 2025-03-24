package main

import (
	"fmt"
	"math"
)

// Define new interface geometry wchih expectes 2 methods area() and perim()
type geometry interface {
	area() float64
	perim() float64
}

// Define 2 structs rect and circle
type rect struct {
	width, height float64
}
type circle struct {
	radius float64
}

// rect can implement geometry interface by implementing ALL methods on the interface
func (r rect) area() float64 {
	return r.width * r.height
}
func (r rect) perim() float64 {
	return 2*r.width + 2*r.height
}

// same for circle
// Note that the methods within each type can be different as can be seen comparing the methods on rect and circle
func (c circle) area() float64 {
	return math.Pi * c.radius * c.radius
}
func (c circle) perim() float64 {
	return 2 * math.Pi * c.radius
}

func measure(g geometry) {
	fmt.Println(g)
	fmt.Println(g.area())
	fmt.Println(g.perim())
}

func detectCircle(g geometry) {
	if c, ok := g.(circle); ok {
		fmt.Println("circle with radius", c.radius)
	}
}

func main() {
	r := rect{width: 3, height: 4}
	c := circle{radius: 5}

	measure(r)
	measure(c)

	detectCircle(r)
	detectCircle(c)
}
