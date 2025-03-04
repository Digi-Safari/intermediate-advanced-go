package main

import (
	"fmt"
	"log"
)

type user struct {
	name  string
	email string
}

func (u user) Write(p []byte) (n int, err error) {
	fmt.Printf("sending notification to %s %s\n", u.name, u.email)
	return len(p), nil
}

func main() {
	u := user{
		name:  "Ray",
		email: "ray@gmail.com",
	}
	//log.LstdFlags

	// end goal is this program compiles
	// user struct variable is passed to log.New
	l := log.New(u, "log: ", log.LstdFlags)
	l.Println("Hello, log")
}
