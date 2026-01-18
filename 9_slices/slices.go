package main

import (
	"fmt"
	"slices"
)

// slice -> Dynamic Arrays -> Most used construct in go -> provides useful methods to add remove etc
func main() {

	// Uninitialized slice is nil (nil = null with respect to other langs)
	// var num []int

	// fmt.Println(num == nil) // true

	// var nums = make([]int, 2, 5) // 2 -> number of elements at initial time, 5 -> capacity (dynamically increases)

	//! 1st method of making slice
	var nums = make([]int, 0, 5) // 2 -> number of elements at initial time, 5 -> capacity (dynamically increases)

	//! 2nd method of making slice
	// nums := []int{}

	// Add in slice
	nums = append(nums, 1)
	nums = append(nums, 2)
	nums = append(nums, 3)
	nums = append(nums, 4)
	nums = append(nums, 5)
	nums = append(nums, 6)

	fmt.Println(nums)
	fmt.Println(len(nums))
	fmt.Println(cap(nums)) // Capacity -> Max numbers of elements can fit

	//! Copy Function
	var copyNums = make([]int, len(nums))

	copy(copyNums, nums) // destination , src

	fmt.Println(copyNums)
	fmt.Println(len(copyNums))
	fmt.Println(cap(copyNums))

	// Slice Operator
	var nums1 = []int{1, 2, 3}

	fmt.Println(nums1[0:1]) // from - to index -> 0 - 1 (Excludes to(1))
	fmt.Println(nums1[1:])
	fmt.Println(nums1[:1])

	// Slice Package - Inbuilt Package
	var n1 = []int{1, 2, 2}
	var n2 = []int{1, 2}

	fmt.Println(slices.Equal(n1, n2)) // return boolean

	// 2D Slices
	var num2d = [][]int{{1, 2, 3}, {3, 4, 5}}
	fmt.Println(num2d)

}
