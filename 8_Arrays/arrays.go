package main

import "fmt"

func main() {
	var nums [4]int
	//array length
	fmt.Println(len(nums))

	nums[0] = 4
	nums[1] = 3
	nums[2] = 2
	nums[3] = 1

	fmt.Println(nums)

	var vals [3]bool
	fmt.Println(vals)

	num := [3]int{1, 2, 3}
	fmt.Println(num)

	//multidimensional array
	var multi [2][3]int
	multi[0][0] = 1
	multi[0][1] = 2

	fmt.Println(multi)

	hello := [3][3]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}

	fmt.Println(hello)

	// Fixed size this is preditable
	// Memory optimization
	// constant time access
	// cache friendly
}
