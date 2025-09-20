package main

import "fmt"

func main() {
	//slices are built on top of arrays

	//dynamic size
	//easy to use
	//built in functions
	//more common than arrays
	//zero value is nil
	//no need to define size

	//uninitialized slice is nil
	var nums []int

	fmt.Println(nums)

	fmt.Println(len(nums))

	var num = make([]int, 3, 5) //intial length, capacity = Maximum numbers of elements slice can hold - can also change like it is also dynamic
	fmt.Println(num)
	fmt.Println(len(num))

	num = append(num, 1)
	num = append(num, 1)
	num = append(num, 1)
	num = append(num, 1)
	num[0] = 4
	fmt.Println(num)
	fmt.Println((cap(num)))

	n := []int{}
	fmt.Println(n)
	fmt.Println(len(n))
	fmt.Println(cap(n))

	//copy function
	a := []int{1, 2, 3}
	b := make([]int, 2)
	copy(b, a)
}
