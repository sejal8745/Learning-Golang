package main

import (
	"fmt"
	"time"
)

func main() {
	//simple switch
	i := 5
	switch i {
	case 1:
		fmt.Println("one")
	case 2:
		fmt.Println("two")
	case 3:
		fmt.Println("three")
	case 4:
		fmt.Println("four")
	case 5:
		fmt.Println("five")
	default: //optional
		fmt.Println("unknown number")
	}

	//switch with multiple expressions
	switch time.Now().Weekday() {
	case time.Saturday, time.Sunday:
		fmt.Println("It's the weekend")
	default:
		fmt.Println("It's a weekday")
	}

	//type switch
	// interface: any type
	whoAmI := func(i interface{}) {

		switch i.(type) {
		// switch t := i.(type) {
		case int:
			fmt.Printf("I'm an int")
		case string:
			fmt.Printf("I'm a string")
		case bool:
			fmt.Printf("I'm a bool")
		default:
			fmt.Printf("Other")
		}

	}

	whoAmI(true)
}
