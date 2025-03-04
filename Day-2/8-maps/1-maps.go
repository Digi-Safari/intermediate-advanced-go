package main

import "fmt"

func main() {
	// keys in map // anytype which can be compared using == can be used as map key
	// value could be of any valid type
	// maps are not concurrency safe

	// provide a len if you already know an estimate of how many values are going to be store in the map

	// by default map is nil // map is a pointer under the hood
	// var m map[string]int  //
	dictionary := make(map[string]string)
	dictionary["up"] = "above"
	dictionary["below"] = "down"

	// Note:- ranging over a map doesn't return values in the same order
	fmt.Println(dictionary)

	for key, value := range dictionary {
		fmt.Println(key, value)
	}
	fmt.Println(len(dictionary))

	delete(dictionary, "up")
	fmt.Println(dictionary)
	modifyMap(dictionary)
	fmt.Println(dictionary)

}

func modifyMap(m map[string]string) {
	m["below"] = "abc"
}
