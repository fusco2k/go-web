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
		cookie, err := request.Cookie("counter")
		if err == http.ErrNoCookie {
			cookie = &http.Cookie{
				Name:  "counter",
				Value: "1",
			}
		}
		counter, err := strconv.Atoi(cookie.Value)
		if err != nil {
			log.Fatalln(err)
		}
		counter++
		cookie.Value = strconv.Itoa(counter)
		http.SetCookie(writer, cookie)
		_, err = io.WriteString(writer, cookie.Value)
		if err != nil {
			log.Fatalln(err)
		}

	})

	log.Fatal(http.ListenAndServe(":8080", nil))
}
