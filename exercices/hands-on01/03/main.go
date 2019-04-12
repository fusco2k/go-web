package main

import (
	"html/template"
	"log"
	"os"
)

type Restaurant struct {
	Name string
	Menu []menu
}

type menu struct {
	Name  string
	Items []item
}

type item struct {
	Name  string
	Price int
}

var tpl *template.Template

func init() {
	tpl = template.Must(template.New("").ParseFiles("tpl.gohtml"))
}

func main() {
	cali := []Restaurant{
		{
			Name:"McDonnalds",
			Menu:[]menu{
				menu{
					Name: "Lunch",
					Items:[]item{
						item{
							Name:  "Coke",
							Price: 2,
						},
						item{
							Name:  "pancakes",
							Price: 5,
						},
						item{
							Name:  "soup",
							Price: 15,
						},
					},
				},
				menu{
					Name: "Dinner",
					Items:[]item{
						item{
							Name:  "Coke",
							Price: 2,
						},
						item{
							Name:  "pancakes",
							Price: 5,
						},
						item{
							Name:  "soup",
							Price: 15,
						},
					},
				},
				menu{
					Name: "Breakfast",
					Items:[]item{
						item{
							Name:  "Coke",
							Price: 2,
						},
						item{
							Name:  "pancakes",
							Price: 5,
						},
						item{
							Name:  "soup",
							Price: 15,
						},
					},
				},
			},
		},
		Restaurant{
			Name:"BK",
			Menu:[]menu{
				menu{
					Name: "Lunch",
					Items:[]item{
						item{
							Name:  "Coke",
							Price: 2,
						},
						item{
							Name:  "pancakes",
							Price: 5,
						},
						item{
							Name:  "soup",
							Price: 15,
						},
					},
				},
				menu{
					Name: "Dinner",
					Items:[]item{
						item{
							Name:  "Coke",
							Price: 2,
						},
						item{
							Name:  "pancakes",
							Price: 5,
						},
						item{
							Name:  "soup",
							Price: 15,
						},
					},
				},
				menu{
					Name: "Breakfast",
					Items:[]item{
						item{
							Name:  "Coke",
							Price: 2,
						},
						item{
							Name:  "pancakes",
							Price: 5,
						},
						item{
							Name:  "soup",
							Price: 15,
						},
					},
				},
			},
		},
		Restaurant{
			Name:"Terraco",
			Menu:[]menu{
				menu{
					Name: "Lunch",
					Items:[]item{
						item{
							Name:  "Coke",
							Price: 2,
						},
						item{
							Name:  "pancakes",
							Price: 5,
						},
						item{
							Name:  "soup",
							Price: 15,
						},
					},
				},
				menu{
					Name: "Dinner",
					Items:[]item{
						item{
							Name:  "Coke",
							Price: 2,
						},
						item{
							Name:  "pancakes",
							Price: 5,
						},
						item{
							Name:  "soup",
							Price: 15,
						},
					},
				},
				menu{
					Name: "Breakfast",
					Items:[]item{
						item{
							Name:  "Coke",
							Price: 2,
						},
						item{
							Name:  "pancakes",
							Price: 5,
						},
						item{
							Name:  "soup",
							Price: 15,
						},
					},
				},
			},
		},
	}
	err := tpl.ExecuteTemplate(os.Stdout, "tpl.gohtml", cali)
	if err != nil {
		log.Fatalln(err)
	}
}
