package main

import "fmt"

func main() {
	//anonymous function, a func without a name
	func(a, b int) {
		fmt.Println(a + b)
	}(10, 20) // () this is how we call anonymous func

	f := func(a, b int) {
		fmt.Println(a+b, "func stored in variable f")
	}
	f(10, 20)

}
