package main

import (
	"fmt"

	"github.com/fatih/color"
	"github.com/nitin0409sep/Go/auth"
	"github.com/nitin0409sep/Go/user" //! Os to has one user package, so import it carefully
)

func main() {
	auth.LoginWithCredentials("nitin0409sep", "secret")

	// auth.extractSession() this will give you error -> as it's scope is till that(auth) package only

	session := auth.GetSession()
	fmt.Println(session)

	u := user.User{
		Name:  "Nitin",
		Email: "nitin0409sep",
		Age:   22,
	}

	fmt.Println(u.Name, u.Email, u.Age, u.Address)
	color.Green("Hello nitin0409sep")
}
