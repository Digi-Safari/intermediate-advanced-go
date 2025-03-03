package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	greet()

	// some further things to do
	// if panic happens before this , it would crash the program and it would stop the further exec
	fmt.Println("end of program")
}

func greet() {
	//fmt.Println(os.Args)
	data := os.Args[1:]
	if len(data) != 3 {
		log.Println("please provide name , age and marks")
		return
	}
	name := data[0]
	ageString := data[1]
	marksString := data[2]
	// errors are values in Go
	// error has default value of nil, which means no error
	//fmt.Println(age, err)

	//if you are calling a function , and if that func returns an error, next thing must be error handling ,
	// you should not continue to write further logic
	age, err := strconv.Atoi(ageString)
	if err != nil {
		log.Println("please provide valid age", err)
		return
	}

	marks, err := strconv.Atoi(marksString)
	if err != nil {
		log.Println("please provide valid marks", err)
		return
	}

	fmt.Println(name, age, marks)

}

//func handleError(err error) {
//	if err != nil {
//		panic(err)
//	}
//}
