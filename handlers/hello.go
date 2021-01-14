// refactoring functionalities into separate handlers; this is one of them

package handlers

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type Hello struct {
	l *log.Logger
}

/*
instead of having a logger object, we make a new 'constructor'
that takes in a specific kind of logger. This 'dependency injection'
allows our code to be more modular and have better unit tests
*/
func NewHello(l *log.Logger) *Hello {
	return &Hello{l}
}

func (h *Hello) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	// use the l object as defined by the constructor above
	h.l.Println("Hello World")

	d, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(rw, "oops!", http.StatusBadRequest)
		return
	}

	fmt.Fprintf(rw, "hello %s\n", d)
}
