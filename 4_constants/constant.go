package main

import "fmt"

const number = "890"

// userName := "user" // Not allowed
var userName string = "user" // Allowed

func main() {

	// Group Constant
	const (
		port = 8000
		host = "localhost"
	)

	fmt.Println(port, host)

	// Single Constant
	const name = "nitin"
	var age = 10

	// name = "tintin"; // Not allowed - you can't reassign it
	age = 20

	fmt.Println(name, age, 890, userName)
}
