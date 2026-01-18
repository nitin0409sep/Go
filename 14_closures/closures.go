package main

import "fmt"

func clouser() func() int {
	var count int = 0

	return func() int {
		count += 1
		return count
	}

}

func main() {
	var increment func() int
	increment = clouser()
	fmt.Println(increment())
	fmt.Println(increment())
}
