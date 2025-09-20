package main

import "fmt"

// in go, whenver the argument is passed to a function, it is passed by value
// this means that a copy of the argument is made and passed to the function
// so any changes made to the argument inside the function do not affect the original variable

// changeNumber tries to change the value of num
func changeNumber(num int) {
	num = 10
	fmt.Println("Inside changeNumber:", num)
}

// changeNumberPointer changes the value of num using a pointer
func changeNumberPointer(num *int) {
	*num = 10 //*num dereferences the pointer to get the value at the address
	fmt.Println("Inside changeNumberPointer:", *num)
}

func main() {
	num := 0
	fmt.Println("Before changeNumber:", num)
	changeNumber(num)
	fmt.Println("After changeNumber:", num)

	fmt.Println("Before changeNumberPointer:", num)
	changeNumberPointer(&num) // &num gets the address of num
	fmt.Println("After changeNumberPointer:", num)

	fmt.Println(("num address:"), &num) // &num gets the address of num
}
