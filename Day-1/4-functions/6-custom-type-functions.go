package main

import (
	"fmt"
	"strings"
)

// creating a new type money, underlying type for money is int
type money int

type stringOps func(str string) string

var dollar money = 100

func main() {
	//time.Duration()
	//time.Second
	fmt.Println(StringManipulation(trimSpace, " This needs to be trimmed "))
	fmt.Println(StringManipulation(toUpper, "This news to be upper"))
	fmt.Println(StringManipulation(greet, "NIC"))
}

func StringManipulation(operation stringOps, str string) string {
	return operation(str)
}

func trimSpace(str string) string {
	return strings.TrimSpace(str)
}

func toUpper(str string) string {
	return strings.ToUpper(str)
}

func greet(str string) string {
	return "Hello, " + str
}
