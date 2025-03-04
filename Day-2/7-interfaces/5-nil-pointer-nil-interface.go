package main

import (
	"database/sql"
	"fmt"
	"log"
)

type DB interface {
	ReadAll()
}

type conn struct {
	db   *sql.DB // default value would be nil
	data string
}

func (c *conn) ReadAll() {
	if c == nil {
		log.Println("c can't be nil")
		return
	}

	c.data = "some data"
	fmt.Println("read all the values from the db", c.data)
}

func main() {
	//var c conn // conn{db:nil,data:""}
	//var c *conn // nil, no memory allocated, accessing any fields would cause panic situation
	c := &conn{} // allocated memory and we have took a reference to that location
	// c := new(conn) // same operation like c := &conn{}

	if c == nil {
		log.Println("c can't be nil")
		return
	}

	var db DB = c
	//fmt.Printf("%#v\n", db)
	// db is not nil // storing a conn type in it
	// if interface is holding any concrete type then it is not nil, even the type value itself could be nil
	// nil interface means no concrete type is present inside the interface
	if db == nil {
		log.Println("db can't be nil")
		return
	}

	db.ReadAll()
}
