package main

import (
	"fmt"
	"log"
	"net/http"
)

type handler int

//handling request and response
func (h handler) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	req.ParseForm()
	//just printing the METHOD request after ParseForm
	_, err := fmt.Fprintf(w, "Message %s", req.Method)
	if err != nil {
		log.Fatalln(err)
	}
}

func main() {
	var h handler
	//enabling listener and connection
	err := http.ListenAndServe(":8080", h)
	if err != nil {
		log.Fatalln(err)
	}
}
