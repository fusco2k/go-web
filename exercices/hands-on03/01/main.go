package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {

	http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {

	})
	http.HandleFunc("/me/", func(writer http.ResponseWriter, request *http.Request) {
		fmt.Fprintln(writer, "Filipe Fusco")
	})
	http.HandleFunc("/dog/", func(writer http.ResponseWriter, request *http.Request) {

	})

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatalln(err)
	}
}
