package main

import (
	"context"
	"fmt"
	"github.com/golang-jwt/jwt/v5"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"rest-api/auth"
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
	// =========================================================================
	// Initialize authentication support
	slog.Info("main : Started : Initializing authentication support")

	publicPEM, err := os.ReadFile("pubkey.pem")
	if err != nil {
		return fmt.Errorf("reading auth public key %w", err)
	}

	publicKey, err := jwt.ParseRSAPublicKeyFromPEM(publicPEM)
	if err != nil {
		return fmt.Errorf("parsing auth public key %w", err)
	}

	a, err := auth.NewAuth(publicKey)
	if err != nil {
		return err
	}
	_ = a // this needs to be changed
	h, err := handlers.API(a)
	if err != nil {
		return err
	}
	api := http.Server{
		Addr:         ":8080",
		ReadTimeout:  8000 * time.Second,
		WriteTimeout: 800 * time.Second,
		IdleTimeout:  800 * time.Second,
		Handler:      h,
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
