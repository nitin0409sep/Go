package main

import "fmt"

// Enums - Enumerated Types
// type OrderStatus int
type OrderStatus string // Custom Type

// const (
//
//	Received OrderStatus = iota  // increments automatically
//	Confirmed
//	Prepared
//	Delivered
//
// )
const (
	Received  OrderStatus = "received"
	Confirmed OrderStatus = "confirmed"
	Prepared  OrderStatus = "prepared"
	Delivered OrderStatus = "delivered"
)

func changeOrderStatus(status OrderStatus) {
	fmt.Println("Changing order status to", status)
}

func main() {
	changeOrderStatus(Confirmed)
}
