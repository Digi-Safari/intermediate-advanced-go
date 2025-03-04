package main

import (
	"fmt"
	"learn-go/Day-2/6-structs/database"
)

func main() {
	conf := database.New("mysql")
	fmt.Println(conf)

	fmt.Println(conf)

	//log.New
	//os.OpenFile()

}
