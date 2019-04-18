package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"net"
	"strings"
)

var s string

func main() {
	ln, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatalln(err)
	}
	defer ln.Close()

	for {
		conn, err := ln.Accept()
		if err != nil {
			log.Fatalln(err)
		}
		go handle(conn)
	}
}

func handle(conn net.Conn) {
	read(conn)
}

func read(conn net.Conn) {
	scan := bufio.NewScanner(conn)
	var i int
	var rMethod, rURI string
	for scan.Scan() {
		ln := scan.Text()
		if i == 0 {
			xs := strings.Fields(ln)
			rMethod = xs[0]
			rURI = xs[1]
			s = rMethod
			fmt.Println("Method:", rMethod)
			fmt.Println("URI:", rURI)
		}
		if ln == "" {
			fmt.Println("end of the request")
			break
		}
		i++
		fmt.Println(ln)
	}
	defer conn.Close()
	write(conn)
}

func write(conn net.Conn) {
	body := `<!DOCTYPE html>
	<html lang="en">
	<head>
	<meta charset="UTF-8">
	<title>Title</title>
	</head>
	<body>
	<p06>"HOLY COW THIS IS LOW LEVEL"</p06>
	</body>
	</html>`

	_, err := io.WriteString(conn, "HTTP/1.1 200 OK\r\n")
	if err != nil {
		log.Fatalln(err)
	}
	_, err = fmt.Fprintf(conn, "Content-Length: %d\r\n", len(body))
	if err != nil {
		log.Fatalln(err)
	}
	_, err = fmt.Fprint(conn, "Content-Type: text/html\r\n")
	if err != nil {
		log.Fatalln(err)
	}
	_, err = io.WriteString(conn, "\r\n")
	if err != nil {
		log.Fatalln(err)
	}
	_, err = io.WriteString(conn, body)
	if err != nil {
		log.Fatalln(err)
	}
	err = conn.Close()
	if err != nil {
		log.Fatalln(err)
	}
}
