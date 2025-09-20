package main

import "fmt"

//iterative purpose

func main() {

	nums := []int{2, 3, 4, 5}

	// normal for loop
	for i := 0; i < len(nums); i++ {
		fmt.Println(nums[i])
	}

	sum := 0
	//range loop
	for _, num := range nums { // _ is idx
		// fmt.Println(num)
		sum += num
	}

	for i, num := range nums {
		fmt.Printf("idx: %d, value: %d\n", i, num)
	}
	fmt.Println("sum:", sum)

	//range on map
	kvs := map[string]string{"a": "apple", "b": "banana"}
	for k, v := range kvs {
		fmt.Printf("%s -> %s\n", k, v)
	}

	//range on string
	//unicode values is the 'c'
	//rune is a datastructure
	// i is the starting byte index of the rune in the string
	for i, c := range "golang" {
		fmt.Println(i, c, string(c)) //c is rune
	}
}
