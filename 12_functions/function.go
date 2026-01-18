package main

import "fmt"

//! Functions can return multiple values

//! Functions - Are 1st class citizens that means we can assign them to a variable or pass/return a func from a func or to a func

// ! Returns 1 value
func add(num1 int, num2 int) int {
	return num1 + num2
}

// ! Returns multiple values
func getLanguages() (string, string, bool) { //! in second written bracket we describe the value type we will be getting in return
	return "go", "js", true
}

// ! Function taking another function as an argument
func processIt(fn func(a int) int) { // func can take a func and also arg func can take a value and return a value
	val := fn(2)
	fmt.Println((val))
}

// ! Return a function from a function
func sum() func(a int, b int) int {
	return func(a int, b int) int {
		return a + b
	}
}

func main() {
	// fmt.Println(add(1, 2))

	// fmt.Println(getLanguages())

	// Can destructure them
	lang1, lang2, _ := getLanguages()
	fmt.Println(lang1, lang2)

	fn := func(a int) int { // Anonymous function
		return 10 + a
	}
	// Passing function as an argument
	processIt(fn)

	var s func(int, int) int = sum()
	// s := sum()
	fmt.Println(s(11, 12))

}
