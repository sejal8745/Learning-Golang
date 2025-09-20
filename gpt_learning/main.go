package main

import "fmt"

func main() {
	score := 30

	if score >= 50 {
		fmt.Println("Pass")
	} else {
		fmt.Println("Fail")
	}

	switch {
	case score >= 90:
		fmt.Println("Excellent")
	case score >= 75:
		fmt.Println("Very Good")
	default:
		fmt.Println("Poor")
	}
}
