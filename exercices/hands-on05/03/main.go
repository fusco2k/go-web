package main

import (
	"html/template"
	"log"
	"net/http"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.New("").ParseFiles("templates/index.gohtml"))
}

func main() {
	fs := http.FileServer(http.Dir("public"))

	http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		err := tpl.ExecuteTemplate(writer, "index.gohtml", nil)
		if err != nil {
			log.Fatalln(err)
		}
	})
	http.Handle("/pics/", fs)

	log.Fatal(http.ListenAndServe(":8080", nil))
}
