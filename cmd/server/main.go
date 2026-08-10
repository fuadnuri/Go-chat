package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	config "www.github.com/fuadnuri/Go-chat/configs"
	"www.github.com/fuadnuri/Go-chat/internal/adapter/httpadapter"
)

func main() {

	rts := httpadapter.Routes()

	config := config.Load()
	server := &http.Server{
		Addr:              fmt.Sprintf("%s:%s", config.Server.Host, config.Server.Port),
		Handler:           rts,
		ReadTimeout:       config.Server.ReadTimeout,
		WriteTimeout:      config.Server.WriteTimeout,
		ReadHeaderTimeout: config.Server.ReadHeaderTimeout,
		IdleTimeout:       config.Server.IdleTimeout,
		MaxHeaderBytes:    1 << 20,
	}
	//start the server
	go func() {
		log.Println("Starting server on port 8000")
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("Server error: %v", err)

		}
	}()
	// waiting for terminal signal
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGINT, syscall.SIGTERM)

	<-stop
	log.Println("Shutting down server...")

	//give existing requests time to complete
	ctx, cancel := context.WithTimeout(
		context.Background(),
		10*time.Second,
	)
	defer cancel()
	if err := server.Shutdown(ctx); err != nil {
		log.Fatalf("Graceful shutdown failed: %v", err)
		//false shutdown remaining connections
		if err := server.Shutdown(context.Background()); err != nil {
			log.Fatalf("Force shutdown failed: %v", err)
		}
	}
	log.Println("Server gracefully stopped")

	//wait for interrupt signal

}
