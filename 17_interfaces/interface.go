package main

import "fmt"

// Interface - r is added as convection
type paymentr interface {
	pay(amount float32)
}

// Razorpay Payment Gateway
type razorpay struct{}

func (r razorpay) pay(amount float32) {
	// Logic to make payment
	fmt.Println("Make payment using razorpay", amount)
}

// Stripe Payment Gateway
type stripe struct{}

func (s stripe) pay(amount float32) {
	// Logic to make payment
	fmt.Println("Make payment using stripe", amount)
}

// Payment Gateway
type payment struct {
	// gateWay razorpay
	gateWay paymentr
}

func (p payment) makePayment(amount float32) {
	// Razorpay Payment
	// razorpayGw := razorpay{}
	// razorpayGw.pay(amount)

	// Stripe Payment
	// stripeGw := stripe{}
	// stripeGw.pay(amount)

	// Ant type of gateway can be passed here now
	p.gateWay.pay(amount)
}

func main() {
	paymentGw := stripe{}

	p1 := payment{
		gateWay: paymentGw,
	}

	p1.makePayment(199)
}
