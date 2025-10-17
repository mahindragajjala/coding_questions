package main

import "fmt"

func main() {
	// Defining and initializing an anonymous struct
	person := struct {
		Name string
		Age  int
	}{
		Name: "Mahindra",
		Age:  28,
	}

	// Accessing fields of the anonymous struct
	fmt.Println("Name:", person.Name)
	fmt.Println("Age:", person.Age)

	// You can also create a slice of anonymous structs
	people := []struct {
		Name string
		Age  int
	}{
		{"Alice", 25},
		{"Bob", 30},
	}

	fmt.Println("\nPeople list:")
	for _, p := range people {
		fmt.Println(p.Name, "-", p.Age)
	}
}
