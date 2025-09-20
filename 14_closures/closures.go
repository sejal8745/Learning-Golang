package main

import "fmt"

// counter function with no argumnet but return a function that return an int
func counter() func() int {
	count := 0
	return func() int {
		count += 1
		return count
	}
}

func main() {
	// Closures are functions that reference variables from outside their body.
	// The function returned by counter is a closure that references the count variable.
	// Each call to counter creates a new instance of count, so multiple counters maintain separate counts.

	// Create a new counter instance
	increment := counter()
	fmt.Println(increment()) // 1
	fmt.Println(increment()) // 2
	fmt.Println(increment()) // 3
}
