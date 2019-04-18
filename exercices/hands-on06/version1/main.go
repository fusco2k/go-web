package main

import (
	"io"
	"log"
	"net/http"
	"strconv"
)

func main() {
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		//check if it has a cookie, if not, creates the cookie counter, if it does, increment the counter
		c, err := request.Cookie("counter")
		if err != nil {
			log.Println(err)
		}
		if c == nil {
			http.SetCookie(writer, &http.Cookie{
				Name:  "counter",
				Value: "1",
			})
			_, err := io.WriteString(writer, c.Value)
			if err != nil {
				log.Fatalln(err)
			}
		} else {
			number, err := strconv.Atoi(c.Value)
			if err != nil {
				log.Fatalln(err)
			}
			number++
			http.SetCookie(writer, &http.Cookie{
				Name:  "counter",
				Value: strconv.Itoa(number),
			})
			_, err = io.WriteString(writer, c.Value)
			if err != nil {
				log.Fatalln(err)
			}
		}
	})

	log.Fatal(http.ListenAndServe(":8080", nil))
}
