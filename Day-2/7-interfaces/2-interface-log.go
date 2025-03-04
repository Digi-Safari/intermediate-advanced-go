package main

import "log"

type user struct {
	name  string
	email string
}

func main() {
	var u user
	//log.LstdFlags

	// end goal is this program compiles
	// user struct variable is passed to log.New
	log.New(u)
}
