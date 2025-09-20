package main

func main() {
	age := 20

	if age < 18 {
		println("Minor")
	} else if age >= 18 && age < 65 {
		println("Adult")
	} else {
		println("Senior Citizen")
	}

	// can declare a variable inside if condition
	if name := "Sahil Patlu Kumar"; name == "Sahil Patlu Kumar" {
		println("He is Mine!")
	}

	// Go doesn't have ternary operator (condition ? true : false)
}
