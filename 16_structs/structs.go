package main

import (
	"fmt"
	"time"
)

// Structs - Custom Data Structure
type order struct {
	id        string
	amount    float32
	status    string
	createdAt time.Time // nanoseconds precision
}

func newOrder(id string, amount float32, status string) *order {
	myOrder := order{
		id:     id,
		amount: amount,
		status: status,
	}

	return &myOrder
}

// Receiver Type
func (o *order) changeStatus(status string) {
	o.status = status // Modifying the status by ref, behind the scene struct handle dereferencing, so we don't need to pass * here as *o.status
}

func (o order) getStatus() string {
	return o.status // Here we don't need to pass ref as we are only reading the value, so it will work fine
}

func main() {

	// 2nd method of creating struct - Inline Struct -> If you want to use struct single time then you can create it
	language := struct {
		name   string
		isGood bool
	}{"golang", true}

	fmt.Println(language)

	// var order order =
	myOrder := order{
		id: "1",
		// amount: 50.1,
		status: "received",
	}

	myOrder.changeStatus("delivered")

	// Assign value to our struct any field
	myOrder.createdAt = time.Now()

	// fmt.Println(myOrder.hello) // You can't access the key that is not defined in struct

	fmt.Println(myOrder.amount) // If you haven't provided a value to the key in struct, by default it will take zeroed value -> int 0, float 0, string "", boolean false

	// Get Field
	fmt.Println(myOrder.status)

	fmt.Println(myOrder)

	myOrder1 := newOrder("2", 20.34, "Hello")

	fmt.Println(*myOrder1)       // dereference
	fmt.Println(myOrder1.amount) // dereferencing not needed in this case, struct handles it on it's own

}
