package main

import "fmt"

// Using Generics in Function
func printSlice[T comparable](items []T) {
	for _, item := range items {
		fmt.Println(item)
	}
}

// func printSlice[T int | string](items []T) {
// 	for _, item := range items {
// 		fmt.Println(item)
// 	}
// }

// Using Generics in Struct
type stack[T int | string] struct {
	elements []T
}

func main() {
	nums := []int{1, 2, 3}
	names := []string{"one", "two", "three"}

	printSlice(nums)
	printSlice(names)

	myStack1 := stack[int]{
		elements: []int{1, 2, 3, 4},
	}

	myStack2 := stack[string]{
		elements: []string{"goland", "js"},
	}

	fmt.Println(myStack1)
	fmt.Println(myStack2)

}
