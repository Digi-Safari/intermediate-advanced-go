package main

import (
	"fmt"
	"strconv"
)

func main() {
	//var err error
	fmt.Println(strconv.Atoi("abc"))
	fmt.Println(strconv.Atoi("qwerty"))
	fmt.Println(strconv.ParseInt("123a", 10, 64))
	fmt.Println(strconv.ParseFloat("xyz", 64))
}
