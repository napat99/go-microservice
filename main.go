package main

import (
	"context"
	"log"
	"microservices_demo/handlers"
	"net/http"
	"os"
	"os/signal"
	"time"
)

func main() {
	// create a logger to pass into our handler
	l := log.New(os.Stdout, "product-api", log.LstdFlags)
	// a handler is a function that would trigger when a path is called
	hh := handlers.NewHello(l)
	gh := handlers.NewGoodbye(l)

	// we have to create a 'mux', basically a router for handers
	sm := http.NewServeMux()
	// generic path -> use hello handler
	sm.Handle("/", hh)
	// specific path -> use specific handler
	sm.Handle("/goodbye", gh)

	// manually create a server to handle things like timeouts
	s := &http.Server{
		Addr:         ":9090",
		Handler:      sm,
		IdleTimeout:  120 * time.Second,
		ReadTimeout:  1 * time.Second,
		WriteTimeout: 1 * time.Second,
	}

	// run the server concurrently so grace shutdown is not blocked
	go func() {
		err := s.ListenAndServe()
		if err != nil {
			l.Fatal(err)
		}
	}()

	// graceful shutdown, allow client to finish request first
	sigChan := make(chan os.Signal)

	// notify us if there is an OS interrupt or kill
	signal.Notify(sigChan, os.Interrupt)
	signal.Notify(sigChan, os.Kill)

	sig := <-sigChan
	l.Println("Received terminate, graceful shutdown", sig)

	tc, _ := context.WithTimeout(context.Background(), 30*time.Second)
	s.Shutdown(tc)
}
