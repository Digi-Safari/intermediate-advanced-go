package main

import "fmt"

func main() {
	//var a any
	var i interface{}

	i = "10"
	i = struct {
		a int
	}{}
	i = 10
	i = true

	// use ok variant to avoid panic
	x, ok := i.(int) // type assertion to fetch concrete value from empty interface
	if ok {
		fmt.Println("value is an integer")
		fmt.Println(x)
	}
	display("hello", 10, true)
	s := []any{"book1", "book2", "book3"}
	display("Name", s...) // unpack the slice and send the individual values to the display func
	display("John")

}

// i is a variadic parameter, it can accept any number of args for the i var
// we can't have a parameter after the variadic parameter
// variadic parameter must be the last parameter in the func signature
// variadic parameter is optional, but we can't set default values
func display(name string, i ...interface{}) {
	fmt.Printf("%#v\n", i)
}
