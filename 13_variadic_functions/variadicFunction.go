package main

import (
	"fmt"
)

func sum(nums ...int) int {
	var sum int

	for _, num := range nums {
		sum = num + sum
	}

	return sum
}

func main() {
	fmt.Println(1, 2, true, "no", 1212, "SA") // Its called as variadic func -> The func in which we can pass n number of args

	fmt.Println(sum(1, 2, 3, 4, 5, 6))

	nums := []int{1, 2, 3, 4}
	fmt.Println(sum(nums...))

}
