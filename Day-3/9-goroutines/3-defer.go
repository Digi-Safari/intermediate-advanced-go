package main

import (
	"fmt"
	"log"
	"os"
)

func main() {

	// defer statement runs when the function is returning
	// defer statements gives guarantee of exec , once registered
	// defer maintains the stack
	// first in last out
	defer fmt.Println(1)
	defer fmt.Println(2)
	//panic("error")
	fmt.Println(3)

	// call the func
	// handle the error
	// the call defer if needed
	f, err := os.Open("file.txt")
	if err != nil {
		log.Println(err)
		return
	}
	defer func() {
		err := f.Close()
		if err != nil {
			log.Println(err)
		}
	}()

}
