package auth

import "fmt"

// If you waana make it public then start name with Capital Letter else if u waana make it private then use small case
func LoginWithCredentials(username string, password string) { // Public Func -> Can be accessed from other modules too
	fmt.Println(username, password)
}
