package main

import (
	"fmt"
)

// Polymorphism means that a piece of code changes its behavior depending on the
// concrete data it’s operating on // Tom Kurtz, Basic inventor

// "Don’t design with interfaces, discover them". - Rob Pike

// Few exceptions- If you want to mock for testing,
//or there are multiple types that need to implement interface immediately

// Bigger the interface weaker the abstraction // Rob Pike

// interface is an abstract type

type Reader interface {

	// interfaces are automatically implemented if method signature is same as of interfaces
	// only method signatures could be added to an interface

	Read(b []byte) (int, error)

	//write(b []byte) (int, error) // all methods must be implemented to implement the interface
}

type File struct {
	Name string
}

func (f File) Read(b []byte) (int, error) {
	fmt.Println("reading files and processing them", f.Name)
	return 0, nil
}

func (f File) PrintFileName() {
	fmt.Println(f.Name)
}

type IO struct {
	name string
}

func (i IO) Read(b []byte) (int, error) {
	fmt.Println("reading and processing io ", i.name)
	return 0, nil
}

func DoReading(r Reader) {

	//piece of code changes its behavior depending on the
	//// concrete data it’s operating on
	// if file is passed Read would be called from the file struct,
	//or if Io is passed Read would be called from IO
	n, err := r.Read(nil)

	//f := r.(File) // type assertion // checking if file object is present in the interface or not
	f, ok := r.(File) // ok would be true if the file object is part of the interface
	if ok {
		f.Name = "file2"
		f.PrintFileName()
	}

	_, _ = n, err
	fmt.Printf("%T\n", r)
}

func main() {

	f := File{Name: "file1"}
	i := IO{name: "io1"}
	DoReading(f)
	DoReading(i)

}
