package main

import "fmt"

func main() {

	i := []int{10, 20, 30, 40, 50}
	b := i[1:3]
	b[0] = 999 // this will change i slice as well
	fmt.Println(b)
	fmt.Println(i)
}

/*
		a:= 10,20,30,40
		b := i[2:5]
a ,(b)-> [10,20,(30,40,50),60] // backing array
// a and b slice shares the same backing array. it is not a copy

b[0] = 999 // b is sharing the same backing array with a slice,
so any updates in b will also result changes in A slice
	    a ,(b)-> [10,20,(999,40,50),60]

*/
