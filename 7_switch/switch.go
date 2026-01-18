package main

import (
	"fmt"
	"time"
)

func main() {
	age := 10

	age = 1

	// simple switch
	switch age {
	case 1:
		{
			fmt.Println("Age", age)
		}
	case 2:
		{
			fmt.Println("Name", age)
		}
	default:
		{
			fmt.Println("Default")
		}
	}

	// Multiple Condition Switch
	switch time.Now().Weekday() {
	case time.Saturday, time.Sunday:
		{
			fmt.Println("Its a weekend")
		}
	default:
		{
			fmt.Println("It's not a weekend")
		}
	}

	// Type Switch
	whatType := func(i interface{}) interface{} { // interface {} => its just like any type
		switch v := i.(type) {
		case int:
			return v
		case string:
			return "string " + v
		default:
			return "none"
		}
	}

	fmt.Println(whatType(10))
	fmt.Println(whatType("nitin"))
	fmt.Println(whatType(90.22))

}
