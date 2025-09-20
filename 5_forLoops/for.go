package main

import "fmt"

// for -> only construct in golang for looping
func main() {

	//while loop implementation using for
	i := 1
	for i <= 4 {
		fmt.Println(i)
		i++
	}

	//infinite loop
	// for {
	// 	fmt.Println("infinite loop")
	// }

	// classic for loop
	for j := 4; j >= 0; j-- {
		// break // to break the loop
		// continue // to skip the current iteration
		if j == 2 {
			continue
		}
		fmt.Println(j)
	}

	//Range loop -> used to iterate over arrays, slices, maps, strings
	for k := range 10 {
		fmt.Println(k)
	}
}
