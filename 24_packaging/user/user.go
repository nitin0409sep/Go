package user

type User struct {
	Age     int
	Name    string
	Email   string
	Address Address
}

// type User struct { // Your struct is public but all key's are private as they are starting from small letter, they can only be accessed in user package
// 	age     int
// 	name    string
// 	email   string
// 	address address
// }

type Address struct {
	Address1 string
}
