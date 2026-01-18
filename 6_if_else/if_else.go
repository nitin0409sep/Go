package main

import "fmt"

func main() {

	age := 20

	//! If  Else
	// if age >= 18 {
	// 	fmt.Println("Person is an adult")
	// } else {
	// 	fmt.Println("Person is not an adult")
	// }

	//! If Else If
	if age > 25 {
		fmt.Println("Person is an adult")
	} else if age > 18 && age < 25 {
		fmt.Println("Person is an adult but less than 25")
	} else {
		fmt.Println("Person is not an adult")
	}

	//! We can declare var inside if
	if name := "abc"; name == "nitiN" {
		fmt.Println(name)
	} else if name == "abc" {
		fmt.Println("abc")
	} else {
		fmt.Println("Nope name doesn't matched")
	}

	//! Go doesn't have ternary operator - you will have to use if else

}
