package main

import "fmt"

// const age int = 30       // also allowed
// var name string = "John" // also allowed

// name := "John" // not allowed

func main() {
	const pi float32 = 3.14

	// pi = 3.15 // throws error: cannot assign to pi (declared const)

	//Multiple constants
	const (
		country    = "USA"
		city       = "New York"
		population = 8419600
	)

	fmt.Println("Country:", country)
}
