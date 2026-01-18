package main

import (
	"fmt"
	"time"
)

type customer struct {
	name   string
	number string
}

type orders struct {
	id        string
	amount    float32
	status    string
	createdAt time.Time

	// Struct Embedding -> Embeded another struct in orders struct
	customer
}

func initializeOrders(id string, amount float32, status string, name string, number string) *orders {
	newOrder1 := orders{
		id:     id,
		amount: amount,
		status: status,
		customer: customer{
			name:   name,
			number: number,
		},
	}

	return &newOrder1
}

func main() {

	order1 := initializeOrders("1", 22.33, "Received", "Nitin", "12112")

	fmt.Println(*order1)

}
