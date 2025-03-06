package main

import (
	"encoding/json"
	"net/http"
)

func main() {

	http.HandleFunc("/home", Home)
	http.HandleFunc("/json", sendJsonResponse)
	// if handler is nil
	// it would use default route matcher available from standard lib, also known as DefaultServeMux
	// http.ListenAndServe blocks forever until its killed
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		panic(err)
	}
}

func Home(w http.ResponseWriter, r *http.Request) {
	//w http.ResponseWriter, is used to write resp to the client
	// http.Request// anything user send us would be in the request struct
	//fmt.Fprintln(w, "Home")
	w.Write([]byte("Hello World"))

}

func sendJsonResponse(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var user struct {
		FirstName string `json:"first_name"` // field tag // we are giving what the name of field should be in json output
		Password  string `json:"-"`          // ignoring the field in JSON output
		Email     string `json:"email"`
	}
	user.FirstName = "abc"
	user.Password = "123"
	user.Email = "abc@gmail.com"

	// json.Marshal converts a type to json
	jsonData, err := json.Marshal(user)
	if err != nil {
		// sending text based error resp
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	// setting header and then sending a response
	w.WriteHeader(http.StatusOK)
	w.Write(jsonData)

}
