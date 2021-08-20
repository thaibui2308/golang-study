package main

import (
	"fmt"
	"net/http"
	"os"
	"os/signal"

	"github.com/braintree/manners"
)

func main() {
	handler := newHandler()

	ch := make(chan os.Signal) // Send signals to this channel when the server is interrupted or killed
	signal.Notify(ch, os.Interrupt, os.Kill)

	go listenForShutdown(ch)

	manners.ListenAndServe(":3010", handler)
}

func newHandler() *handler {
	return &handler{}
}

type handler struct{}

func (h *handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()

	name := query.Get("name")
	if name == "" {
		name = "Thai Bui"
	}

	fmt.Fprintf(w, "Hello, my name is ", name)
}
func listenForShutdown(ch <-chan os.Signal) {
	<-ch
	manners.Close()
}
