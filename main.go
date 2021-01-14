package main

import (
	"log"
	"microservices_demo/handlers"
	"net/http"
	"os"
)

func main() {
	// create a logger to pass into our handler
	l := log.New(os.Stdout, "product-api", log.LstdFlags)
	// a handler is a function that would trigger when a path is called
	hh := handlers.NewHello(l)

	// we have to create a 'mux', basically a router for handers
	sm := http.NewServeMux()
	sm.Handle("/", hh)

	// specify a port for server and attach mux to it, nil means just use default mux
	http.ListenAndServe(":9090", sm)

}
