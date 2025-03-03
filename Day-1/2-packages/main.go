package main

import (
	"fmt"
	"learn-go/Day-1/db"
	"learn-go/Day-1/sum"
)

func main() {

	sum.Add(10, 20)
	sum.Sub(100, 90)
	fmt.Println()

	sum.Multiply(10, 10)
	sum.MultiplyResult = 200
	fmt.Println(sum.MultiplyResult)

	// don't use exported global variables
	// if you don't want the results to be modified by others

	// constants in Global scope are perfectly fine, they can't be modified
	db.PostgresConn = "mysql"
	fmt.Println(db.PostgresConn)
}
