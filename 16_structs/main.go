package main

import (
	"fmt"
	"time"
)

//order struct

type order struct {
	id        int
	amount    float64
	status    string
	createdAt time.Time // nano sec precision ~accurate
}

// constructors - function that returns struct instance
func newOrder(id int, amount float64, status string) *order {
	return &order{
		id:        id,
		amount:    amount,
		status:    status,
		createdAt: time.Now(),
	}
}

// struct methods - receiver type
// func changeStatus(status string) {} ------- normal function
func (o *order) changeStatus(status string) {
	o.status = status

	// no need of deferecing unline normal function. --- *o.status = status
	// fmt.Println("status changed to: ", o.status)
}

func (o *order) getAmount() float64 {

	// since here we aren;t not changing the values hence we can also remove the * in function syntax (o order)
	return o.amount
}

func main() {
	// struct is a collection of fields
	// fields are variables that belong to the struct
	// fields can be of different types
	//creating instance

	// If we don't set the fields, they will be set to their zero values
	// int -> 0, float64 -> 0.0, string -> "", time.Time -> zero time

	// order1 := order{id: 1, amount: 100.50, status: "pending", createdAt: time.Now()} // all fields initialized
	// order2 := order{id: 2} // only id field initialized, rest will be zero values
	// order3 := order{} // all fields will be zero values
	var order1 order = order{
		id:        1,
		amount:    100.50,
		status:    "pending",
		createdAt: time.Now(),
	}

	order2 := order{
		id:        2,
		amount:    200.75,
		status:    "completed",
		createdAt: time.Now(),
	}

	fmt.Println("order1 struct: ", order1)
	fmt.Println("order2 struct: ", order2)

	//get values of fields
	fmt.Println("order1 amount: ", order1.amount)

	order1.changeStatus("shipped")
	fmt.Println("order1 status after changeStatus method call: ", order1.status)

	fmt.Println("order2 amount using getAmount method: ", order2.getAmount())

	// using constructor
	// the & shows up because your constructor returns a pointer (*order), and Go’s default fmt.Println format prints pointers as &{...}.
	order3 := newOrder(3, 300.00, "pending")
	fmt.Println("order3 struct using constructor: ", *order3)
	fmt.Println("order3 struct using constructor with '&': ", order3)

	//inline struct

	languages := struct {
		name    string
		ranking int
	}{
		name:    "Go",
		ranking: 1,
	}

	fmt.Println("inline struct: ", languages) // for one instance usecase

}
