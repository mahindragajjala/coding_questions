package main

import "fmt"

// Basic interfaces
type Area interface {
	Area() float64
}

type Perimeter interface {
	Perimeter() float64
}

// Embedded interface
type Shape interface {
	Area      // embeds Area interface
	Perimeter // embeds Perimeter interface
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

/* func main() {
	var s Shape = Circle{Radius: 5}

	fmt.Printf("Area: %.2f\n", s.Area())
	fmt.Printf("Perimeter: %.2f\n", s.Perimeter())
}
*/

func Constructor_function(c Shape) {
	fmt.Println(c.Area())
	fmt.Println(c.Perimeter())
}
func main() {
	c1 := Circle{
		Radius: 4,
	}
	Constructor_function(c1)
}
