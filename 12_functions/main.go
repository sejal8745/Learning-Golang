package main

import "fmt"

func add(x int, y int) int {
	return x + y
}

func getLanguages() (string, string) {
	return "Go", "Python"
}

func processIt(fn func(a int) int) int {
	return fn(3)
}

// return function
func processIt2() func(a int) int {
	return func(a int) int {
		return a * a
	}
}

func main() {
	result := add(3, 4)
	fmt.Println(result)

	lang1, lang2 := getLanguages()
	fmt.Println(lang1, lang2)

	//if you want to ignore one of the return values
	lang3, _ := getLanguages() // _ is blank identifier
	fmt.Println(lang3)

	fn := func(a int) int {
		return a * a
	}

	fmt.Println(processIt(fn))
	fn1 := processIt2()

	fmt.Println(fn1(4))

}
