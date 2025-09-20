package main

import "fmt"

func main() {
	//string
	var var_name string = "Hello, World!"
	//if the variable is not used, it will throw an error

	// var var_name = "Hello, World!" is also valid:: type inferred
	// var_name := "Hello, World!"
	// is also valid:: type inferred and declaration + assignment together

	// ** when value is not asssigned then var, and type is mandatory
	// var var_name string
	// var_name = "Hello, World!"**

	fmt.Println(var_name)

	//int
	var var_name2 int = 123
	fmt.Println(var_name2)

	//boolean
	var var_name3 bool = true
	fmt.Println(var_name3)

	//float
	var var_name4 float32 = 3.14
	fmt.Println(var_name4)
}
