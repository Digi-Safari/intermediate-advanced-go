package main

import "fmt"

func main() {
	// add would be passed to op parameter, 10 and 20 would be passed to x and y variables
	operate(add, 2, 3)
	operate(sub, 2, 3)
}

// operate func can accept function in op parameter,
// the function signature we are passing should match to op parameter type
func operate(op func(int, int) int, x, y int) {
	sum := op(x, y)
	fmt.Println(sum)
}

// datatype of func -> func(args)returnType
func add(a, b int) int {
	return a + b

}

func sub(a, b int) int {
	return a - b
}
