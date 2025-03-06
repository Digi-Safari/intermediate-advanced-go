package main

import "net/http"

// middleware that exec some pre-processing or the post-processing logic
// req -> mid1->mid-2-> handler->services
//
//	<-		   <-			<-				 <-            <- return flow
//
// Middleware Examples: logging, Panic Recovery, Auth, Authorize, GenerateReqID, Fetching Headers
func main() {
	http.HandleFunc("/home", Mid(home))
}

func home(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Home Page"))
}
func Mid() {

}
