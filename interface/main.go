package main

import "fmt"

// Interface
type Shape interface {
	Area() float64
	Perimeter() float64
}

// Circle struct
type Circle struct {
	Radius float64
}

func (c Circle) Area() float64 {
	return 3.14159 * c.Radius * c.Radius
}
func (c Circle) Perimeter() float64 {
	return 2 * 3.14159 * c.Radius
}

// Rectangle struct
type Rectangle struct {
	Length, Breadth float64
}

func (r Rectangle) Area() float64 {
	return r.Length * r.Breadth
}
func (r Rectangle) Perimeter() float64 {
	return 2 * (r.Length + r.Breadth)
}

// Function using interface
func printShapeInfo(s Shape) {
	fmt.Printf("Area: %.2f, Perimeter: %.2f\n", s.Area(), s.Perimeter())
}

func main() {
	c := Circle{Radius: 5}
	r := Rectangle{Length: 4, Breadth: 6}

	// Using interface
	printShapeInfo(c)
	printShapeInfo(r)
}
