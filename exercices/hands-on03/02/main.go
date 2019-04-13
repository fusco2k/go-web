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

	http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {

	})
	http.HandleFunc("/me", func(writer http.ResponseWriter, request *http.Request) {
		err := tpl.ExecuteTemplate(writer, "tpl.gohtml", "Filipe")
		if err != nil {
			log.Fatalln(err)
		}
	})
	http.HandleFunc("/dog/", func(writer http.ResponseWriter, request *http.Request) {

	})

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatalln(err)
	}
}
