package main

import (
	"fmt"
	"maps"
)

func main() {
	//maps: hash object, dictionary, associative array

	//creating map
	mp := make(map[string]int) //map[keyType]valueType

	//adding key-value pairs
	mp["age"] = 30
	mp["score"] = 100

	//accessing values
	age := mp["age"]
	fmt.Println("Age:", age)

	//imp: if key not present, returns zero value of valueType
	height := mp["height"]         //not present
	fmt.Println("Height:", height) //0

	//length of map
	fmt.Println("Length:", len(mp))

	//deleting key-value pair
	delete(mp, "score")
	fmt.Println("Length after delete:", len(mp), mp)

	//clearing map
	mp = make(map[string]int) //reinitialize
	fmt.Println("Length after clearing:", len(mp), mp)

	//without make
	mp2 := map[string]string{
		"name":    "Alice",
		"country": "Wonderland",
	}

	fmt.Println("Map 2:", mp2)

	//check if element exists
	// mp2["name"] returns two values, second is boolean ----- imp
	value, exists := mp2["name"]
	if exists {
		fmt.Println("Name exists:", value)
	} else {
		fmt.Println("Name does not exist")
	}

	//chrck if two maps are equal
	mp3 := map[string]string{
		"name":    "Alice",
		"country": "Wonderlands",
	}

	//direct comparison not allowed
	// fmt.Println(mp2 == mp3) //compile error

	fmt.Println(maps.Equal(mp2, mp3))

}
