package main

import (
	"fmt"
	"maps"
)

// maps -> hash, object, dict
func main() {

	//! Method 1 - creating map
	m := make(map[string]string) // map[key : type]value : type

	//! Method 2 - creating map --> When you already know elements use this else use 1st method
	m2 := map[string]int{"price": 10}
	m3 := map[string]int{"price": 10, "age": 11}
	fmt.Println((m2))

	// setting an element
	m["name"] = "golang"
	m["key"] = "key"

	// get an element value
	fmt.Println(m["name"]) // Print map value
	fmt.Println(m["key"])  // Print map value

	//! IMP: if key doesn't exist in map -> and we try to access it, it will return us zeroed value i.e -> string = "", boolean = false, number = 0
	fmt.Println(m["hi"]) // Print map value

	//! Length of map
	fmt.Println(len((m)))

	//! Delete element
	delete(m, "key")

	fmt.Println(m)
	fmt.Println(len((m)))

	//! Clear the whole map
	clear(m)
	fmt.Println(m)

	//! Check whether value exists in map or not
	_, ok := m2["price1"] // _ => value of key -> may be written as v or anything -> rn its value will be 0 if not okay else 10

	if ok {
		fmt.Println("Okay")
	} else {
		fmt.Println("Not okay")
	}

	//! Check maps are equal or not
	// fmt.Println(maps.Equal(m, m2));  // This will give error as type of key value is not similar in both

	fmt.Println(maps.Equal(m2, m3))
}
