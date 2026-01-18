package main

import "fmt"

func main() {

	// Situation 1 - U have declared a var but don't wanna assign value rn
	var userAge int
	userAge = 10

	fmt.Println(userAge)

	// Short Hand Syntax
	userName := "Nitin"

	fmt.Println(userName)

	// var name string = "nitin" // If you have declared any var, then its must that you will have to use it otherwise you will have to delete it

	// infers the type
	var name = "golang"
	var isAdult = true
	var num = 1 
	


	fmt.Println(name)
	fmt.Println(isAdult)
	fmt.Println(num)
}
