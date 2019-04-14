package main

import (
	"io"
	"log"
	"net/http"
	"html/template"
)

func main() {

	http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		_, err := io.WriteString(writer, "Foo ran")
		if err != nil {
			log.Fatalln(err)
		}
	})
	http.HandleFunc("/dog/", func(writer http.ResponseWriter, request *http.Request) {
		tpl, err := template.ParseFiles("dog.gohtml")
		if err != nil {
			log.Fatalln(err)
		}
		err = tpl.ExecuteTemplate(writer, "dog.gohtml", nil)
		if err != nil {
			log.Fatalln(err)
		}
	})

	http.Handle("/dog.jpeg", http.FileServer(http.Dir(".")))

	log.Fatal(http.ListenAndServe(":8080", nil))
}
