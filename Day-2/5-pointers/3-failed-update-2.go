package main

import "fmt"

func main() {
	x := 10
	p := &x
	updateVal(p)

	fmt.Println(x)

}

func updateVal(px *int) {

	var randomVal = 200
	// we have changed the reference of px to store a new variable
	px = &randomVal

	// this would not update the main function x variable, it would update randomVal
	*px = 100
}
