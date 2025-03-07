package main

import (
	"errors"
	"fmt"
	"log"
	"strconv"
)

// Error type should end with the word Error
// Error types should not be used for any domain data
// Creating error struct could be useful when dealing with dynamic error data

type QueryError struct {
	Func  string
	Input string
	Err   error
}

// Error() method is implemented to implement error interface
func (q *QueryError) Error() string {
	// formatting the output string
	return "main." + q.Func + ": " + "input " + q.Input + " " + q.Err.Error()
}

func SearchSomething(id int) (string, error) {
	// assume that search code is written and we need to return an error
	return "", &QueryError{
		Func:  "SearchSomething",
		Input: strconv.Itoa(id),
		Err:   errors.New("not found"),
	}
}

func main() {
	_, err := SearchSomething(100)
	if err != nil {

		var qe *QueryError
		// using errors.As we can check if custom struct is present in the chain or not
		ok := errors.As(err, &qe)
		if ok {
			fmt.Printf("errorsAs %#v\n", qe.Input) // we can access individual fields if needed, or take some specific actions
			return
		}
		log.Println(err)
		return
	}

}
