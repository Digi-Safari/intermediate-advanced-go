package main

import "fmt"

func main() {
	//Arrays size is fixed, we can't grow them
	var i [10]bool

	i[0] = true

	b := [5]int{10, 20, 30}
	fmt.Println(i, b)

	c := [...]int{1, 2, 3, 4, 5} //... would create the array size according to the number of values passed
	// if three values are passed, then size would three, but after creation we cant grow the array
	fmt.Println(c)
	//c[6] = 10
	fmt.Println("len of c is ", len(c))

	//for i := 0; i < len(c); i++ {
	//	fmt.Println(c[i])
	//}
	//x := 0
	//for x < len(c) {
	//	fmt.Println(c[x])
	//	x++
	//}

	for index, value := range c {
		fmt.Println(index, value)
	}

	// using _ , we are ignoring the return value from the range
	for _, value := range c {
		fmt.Println(value)
	}
}
