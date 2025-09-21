package main

import "fmt"

type customer struct {
	name  string
	email string
}

type order struct {
	id       int
	amount   float64
	status   string
	customer // embedded struct
}

func main() {

	myOrder := order{
		id:     1,
		amount: 250.75,
		status: "shipped",
	}

	fmt.Println("Order", myOrder) // Order {1 250.75 shipped { }}
	// { } -> zero value of embedded struct

	// setting values of embedded struct
	myOrder.customer = customer{
		name:  "John Doe",
		email: "",
	}
	// here while embidding struct we don't give any name to it (customer customer)
	// so we can access the fields of embedded struct directly like below
	// myOrder.name = "John Doe"
	// myOrder.email = "jon@gmail.com"

	fmt.Println("Order", myOrder) // Order {1 250.75 shipped {John Doe }}
}
