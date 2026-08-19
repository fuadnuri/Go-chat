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

	config "github.com/fuadnuri/Go-chat/configs"
	"github.com/fuadnuri/Go-chat/internal/adapter/httpadapter"
)

var rts = httpadapter.Routes()

func server() *http.Server {
	cfg := config.Load()
	return &http.Server{
		Addr:              fmt.Sprintf("%s:%s", cfg.Server.Host, cfg.Server.Port),
		Handler:           rts,
		ReadTimeout:       cfg.Server.ReadTimeout,
		WriteTimeout:      cfg.Server.WriteTimeout,
		ReadHeaderTimeout: cfg.Server.ReadHeaderTimeout,
		IdleTimeout:       cfg.Server.IdleTimeout,
		MaxHeaderBytes:    1 << 20,
	}
}

func main() {

	server := server()
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
