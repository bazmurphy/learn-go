package main

import "fmt"

func main() {
	user, err := getUser()
	if err != nil {
		fmt.Println(err)
		return
	}
	// use the "user" here

	profile, err := getUserProfile(user.ID)
	// ^this depends on the function above for its argument user.ID
	if err != nil {
		fmt.Println(err)
		return
	}
	// use the "profile" here
}

// by looking at this function signature we can see it returns a User and possibly an error we need to handle
func getUser() (User, err) {
	// do some get user logic here
}

// Go Error handling forces you to be hyper aware of every potential error and that you write code which handles it
