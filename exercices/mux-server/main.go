package main

import (
	"io"
	"net/http"
)

func a(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "Something to display")
}

func main() {

	http.HandleFunc("/", a)

	http.ListenAndServe(":8080", nil)
}
