package main

import "fmt"

type Student struct {
	name string
	age  int
}

// method signature
//func (receiver) funcSignature {//body}

func (s Student) SayHello() {
	fmt.Println("Hello, my name is", s.name)
}

// this is not allowed
//func (s Student) SayHello() string {
//	fmt.Println("Hello, my name is", s.name)
//}

func main() {

	s := Student{
		name: "Jim",
		age:  30,
	}
	// methods could be only called using struct variable
	s.SayHello()
	fmt.Println(s)
}
