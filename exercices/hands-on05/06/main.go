package main

import (
	"html/template"
	"log"
	"net/http"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.New("").ParseGlob("templates/*"))
}

func main() {
	http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		err := tpl.ExecuteTemplate(writer, "index.gohtml", nil)
		handleError(err, writer)
	})

	http.HandleFunc("/about", func(writer http.ResponseWriter, request *http.Request) {
		err := tpl.ExecuteTemplate(writer, "about.gohtml", nil)
		handleError(err, writer)
	})

	http.HandleFunc("/apply", func(writer http.ResponseWriter, request *http.Request) {
		if request.Method == http.MethodPost {
			err := tpl.ExecuteTemplate(writer, "applyProcess.gohtml", nil)
			handleError(err, writer)
			return
		}
		err := tpl.ExecuteTemplate(writer, "apply.gohtml", nil)
		handleError(err, writer)
	})

	http.HandleFunc("/applyProcess", func(writer http.ResponseWriter, request *http.Request) {
		err := tpl.ExecuteTemplate(writer, "applyProcess.gohtml", nil)
		handleError(err, writer)
	})

	http.HandleFunc("/contact", func(writer http.ResponseWriter, request *http.Request) {
		err := tpl.ExecuteTemplate(writer, "contact.gohtml", nil)
		handleError(err, writer)
	})

	log.Fatal(http.ListenAndServe(":8080", nil))
}

func handleError(err error, writer http.ResponseWriter) {
	if err != nil {
		http.Error(writer, err.Error(), http.StatusInternalServerError)
	}
}
