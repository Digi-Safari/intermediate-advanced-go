package main

import (
	"log"
	"net/http"
)

// main-> if request hits -> handlers -> models
func main() {
	// globally set config for log package
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	// it would initialize the conn once at the start of the app
	setupRoutes()
	http.ListenAndServe(":8080", nil)

}
