package main

import (
	"log"
	"os"
	"text/template"
)

type hotel struct {
	Name, Address, City, Zip string
}

type region struct {
	Region string
	Hotels []hotel
}

type Regions []region

var tpl *template.Template

func init() {
	tpl = template.Must(template.New("").ParseFiles("tpl.gohtml"))
}

func main() {
	cali := Regions{
		region{
			Region: "southern",
			Hotels: []hotel{
				hotel{
					Name:    "Atma",
					Address: "Street 123",
					City:    "California",
					Zip:     "12345",
				},
				hotel{
					Name:    "Bayside",
					Address: "Street 123",
					City:    "California",
					Zip:     "12345",
				},
			},
		},
		region{
			Region: "central",
			Hotels: []hotel{
				hotel{
					Name:    "Atma",
					Address: "Street 123",
					City:    "California",
					Zip:     "12345",
				},
				hotel{
					Name:    "Bayside",
					Address: "Street 123",
					City:    "California",
					Zip:     "12345",
				},
			},
		},
		region{
			Region: "northern",
			Hotels: []hotel{
				hotel{
					Name:    "Atma",
					Address: "Street 123",
					City:    "California",
					Zip:     "12345",
				},
				hotel{
					Name:    "Bayside",
					Address: "Street 123",
					City:    "California",
					Zip:     "12345",
				},
			},
		},
	}

	err := tpl.ExecuteTemplate(os.Stdout, "tpl.gohtml", cali)
	if err != nil {
		log.Fatalln(err)
	}
}
