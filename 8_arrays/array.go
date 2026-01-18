package main

import "fmt"

// zeroed values = int -> 0, string -> "", boolean -> false

func main() {

	// number sequenced of specific length
	var nums [4]int // zeroed values -> By default values (called as zeroed values) on all index's - int -> 0, string -> '', boolean -> false

	nums[0] = 1

	fmt.Println(len(nums))
	fmt.Println(nums)

	//! to declare it in single line
	ages := [3]int{18, 20, 22}
	fmt.Println(ages)

	//! 2d Array
	numbers := [2][2]int{{1, 1}, {2, 3}}
	fmt.Println(numbers)

	// When to use array -> Fixed Size, that is predictable
	// Memory Optimization
	// Constant time access

	// if we don't know exact length and all, then we used slices that provide or allocate memory dynamically
}
