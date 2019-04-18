package main

import (
	"bufio"
	"encoding/csv"
	"html/template"
	"log"
	"os"
	"strconv"
)

var tpl *template.Template

type Data struct {
	Date string
	Open float64
}

func init() {
	tpl = template.Must(template.New("").ParseFiles("tpl.gohtml"))
}

func main() {
	var record []Data

	file, err := os.Open("table.csv")
	if err != nil {
		log.Fatalln(err)
	}
	defer file.Close()

	csvOut := csv.NewReader(bufio.NewReader(file))
	output, err := csvOut.ReadAll()

	for i, row := range output {
		if i == 0 {
			continue
		}
		date := row[0]
		open, _ := strconv.ParseFloat(row[1], 64)

		record = append(record, Data{
			Date: date,
			Open: open})
	}

	err = tpl.ExecuteTemplate(os.Stdout, "tpl.gohtml", record)

	if err != nil {
		log.Fatalln(err)
	}
}
