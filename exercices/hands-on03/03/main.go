package main

import (
	"html/template"
	"log"
	"net/http"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.New("").ParseFiles("tpl.gohtml"))
}

func main() {

	http.Handle("/", http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {

	}))

	http.Handle("/me", http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		err := tpl.ExecuteTemplate(writer, "tpl.gohtml", "Filipe")
		if err != nil {
			log.Fatalln(err)
		}
	}))
	http.Handle("/dog/", http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {

	}))

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatalln(err)
	}
}
