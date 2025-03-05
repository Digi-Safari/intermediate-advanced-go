package main

import "fmt"

type user struct {
	name string
	age  int
}

type users map[int]*user

func main() {
	//var u users = make(users)
	u := users{1: &user{"john", 20}, 2: &user{"jane", 21}}

	usr, ok := u[5] // ok would be true if user is found
	if !ok {
		fmt.Println("user not found")
		return
	}
	fmt.Println(usr.name)
}
