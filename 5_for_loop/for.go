package main

// println - Intended mainly for debugging the compiler/runtime

import "fmt"

// for - For looping only for is present in go, no while and do wile is present in go
// By using for only we implement while loop also

func main() {

	// While Loop using for loop
	i := 1

	for i <= 3 {
		fmt.Println((i))
		i++
	}

	//! infinite loop
	// for {
	// 	println("Infinite")
	// }

	// Classic for loop
	for i := 1; i < 6; i++ {

		// U can use continue and break
		if i == 2 {
			continue
		}

		if i == 4 {
			break
		}

		println(i) // println - Intended mainly for debugging the compiler/runtime
	}

	// Range - came in 1.22 go lang
	for i := range 3 {
		fmt.Println((i))
	}

}
