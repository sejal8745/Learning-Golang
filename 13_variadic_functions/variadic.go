package main

import "fmt"

// function which could have n numbers of arguments with may be of different types
// e.g. Println function in fmt package

func sum(nums ...int) int {
	total := 0
	for _, num := range nums {
		total += num
	}
	return total
}

// any type
func printAll(args ...interface{}) {
	for _, arg := range args {
		fmt.Println(arg)
	}
}

func main() {

	result := sum(1, 2, 3, 6, 8)
	fmt.Println(result)

	printAll(1, "hello", 3.4, true)

	// if you have a slice and you want to pass it to a variadic function
	nums := []int{1, 2, 3, 4, 5}
	result2 := sum(nums...) //... is required like a spread operator in js
	fmt.Println(result2)
}
