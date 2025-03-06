package main

import (
	"net/http"
	"small-app/handlers"
)

func setupRoutes() {
	http.HandleFunc("/find", handlers.FindUser)
}
