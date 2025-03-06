package main

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"rest-api/handlers"
	"time"
)

func main() {
	setupSlog()
	err := startApp()
	if err != nil {
		panic(err)
	}
}

func startApp() error {
	api := http.Server{
		Addr:         ":8080",
		ReadTimeout:  8000 * time.Second,
		WriteTimeout: 800 * time.Second,
		IdleTimeout:  800 * time.Second,
		Handler:      handlers.API(),
	}

	// channel to store any errors while setting up the service
	serverErrors := make(chan error)
	go func() {
		serverErrors <- api.ListenAndServe()
	}()

	//shutdown channel intercepts ctrl+c signals
	shutdown := make(chan os.Signal, 1)
	signal.Notify(shutdown, os.Interrupt)

	select {
	case err := <-serverErrors:
		return err
	case <-shutdown:
		fmt.Println("graceful shutdown started")

		//Shutdown gracefully shuts down the server without interrupting any active connections.
		//Shutdown works by first closing all open listeners
		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()
		err := api.Shutdown(ctx)
		if err != nil {
			//Close immediately closes all active net.Listeners
			err := api.Close()
			if err != nil {
				return err
			}
		}

	}

	return nil
}
