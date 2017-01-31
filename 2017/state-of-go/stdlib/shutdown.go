package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
)

func main() {
	// subscribe to SIGINT signals
	quit := make(chan os.Signal) // HL
	signal.Notify(quit, os.Interrupt)

	srv := &http.Server{Addr: ":8080", Handler: http.DefaultServeMux}
	go func() { // HL
		<-quit // HL
		log.Println("Shutting down server...")
		if err := srv.Shutdown(context.Background()); err != nil { // HL
			log.Fatalf("could not shutdown: %v", err)
		}
	}()

	http.HandleFunc("/", handler)
	err := srv.ListenAndServe()
	if err != http.ErrServerClosed { // HL
		log.Fatalf("listen: %s\n", err)
	}
	log.Println("Server gracefully stopped")
}

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "hello")
}
