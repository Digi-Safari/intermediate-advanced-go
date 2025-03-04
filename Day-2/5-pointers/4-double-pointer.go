package main

import (
	"fmt"
)

func main() {
	// assume p address is x80 and value is nil
	var p *int           // nil
	updateNilPointer(&p) // passing address of p (x80)

	// pointer would have updated value because of double pointer
	fmt.Println(*p)

	// both of them uses double pointer concepts
	//json.Unmarshal()
	//errors.As()
}

func updateNilPointer(p1 **int) {
	// assume p1 is storing x80(address of p from the main)
	x := 10 // assume address x90

	// trying to access the value of p1
	// which is also another pointer named as p from the main function
	*p1 = &x // updating x80 = x90 // it directly changes p from the main function itself
}

/*
**p1 -> [*p]
 */
