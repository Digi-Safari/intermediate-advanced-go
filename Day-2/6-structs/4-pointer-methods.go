package main

import (
	"fmt"
)

// when to use a pointer
// https://go.dev/doc/faq#methods_on_values_or_pointers:~:text=Should%20I%20define%20methods%20on%20values%20or%20pointers%3F%C2%B6
type author struct {
	name string
	age  int
}

func (a *author) updateName(name string) {
	a.name = name
}
func main() {
	a := author{name: "zhangsan", age: 18}
	a.updateName("Bob")
	fmt.Println(a.name)

}
