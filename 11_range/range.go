package main

import "fmt"

// Range - Iterating over data structures
func main() {

	// Iterate over slice
	nums := []int{6, 7, 8}
	sum := 0
	for i, num := range nums {
		fmt.Print(i)     // Index
		fmt.Println(num) // Index
		sum = sum + num
	}

	fmt.Println(sum)

	// Iterate over maps
	m := map[string]int{"price": 10, "age": 20}

	for k, v := range m {
		fmt.Println(k, v) // k - key, v - value
	}

	//! Iterate over string
	str := "Nitin"

	//! unicode code point rune
	//! starting byte of rune i.e. i is not index here exactly
	//! 300 -> 1 byte , 2 byte

	for i, c := range str {
		fmt.Println(i, c) // i -> index , c -> Character // This will print ASCII values of character
		fmt.Println(i, string(c))
	}

	// for i := 0; i < len(nums); i++ {
	// 	fmt.Println(nums[i])
	// }

}
