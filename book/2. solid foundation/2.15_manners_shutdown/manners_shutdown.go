package main

import (
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"github.com/braintree/manners"
)

func main()  {
	handler := newHandler()

	// настраиваем мониторинг сигналов ОС
	ch := make(chan os.Signal)
	signal.Notify(ch, os.Interrupt, os.Kill)
	go listenForShutdown(ch)

	// start web-server
	manners.ListenAndServe(":8080", handler)
}

func newHandler() *handler  {
	return &handler{}
}

type handler struct {}

func (h *handler) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	query := req.URL.Query()
	name := query.Get("name")
	if name == "" {
		name = "Inigo Montoya"
	}
	fmt.Fprint(res, "Hello, my name is ", name)
}

func listenForShutdown(ch <-chan os.Signal)  {
	<-ch
	manners.Close()
}