package main

import "fmt"

// Args - pass by value
func changeNum(num int) {
	num = 5
	fmt.Println("Change Num -", num)
}

// Args - pass by Reference
func changeNumByRef(num *int) {
	*num = 5 //! Dereferencing

	fmt.Println("Change Num By Ref -", *num)
}

func main() {
	num := 1

	changeNum(num) // Passed as a value

	fmt.Println("After change num -", num)

	changeNumByRef(&num) // Passed as a reference

	fmt.Println("After change num -", num)
}
