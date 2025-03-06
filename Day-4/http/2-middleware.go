package main

import (
	"fmt"
	"net/http"
)

// middleware that exec some pre-processing or the post-processing logic
// req -> mid1->mid-2-> handler->services
//
//	<-		   <-			<-				 <-            <- return flow
//
// Middleware Examples: logging, Panic Recovery, Auth, Authorize, GenerateReqID, Fetching Headers
func main() {

	http.HandleFunc("/home", Mid(Mid2(home)))
	panic(http.ListenAndServe(":8080", nil))

}

func home(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Home Page")
	w.Write([]byte("Home Page"))
}
func Mid(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Mid layer started")
		fmt.Println("pre processing logic")
		next(w, r)
		fmt.Println("post processing logic")
		fmt.Println("Mid layer ended")
	}

}

func Mid2(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Mid2 layer started")
		next(w, r)
		fmt.Println("Mid2 layer ended")
	}
}
