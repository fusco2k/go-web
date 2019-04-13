package main

import (
	"io"
	"log"
	"net"
)

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

func handle(conn net.Conn){
	_, err:=io.WriteString(conn, "I see you connected\n")
	if err != nil {
		log.Fatalln(err)
	}
	err = conn.Close()
	if err != nil {
		log.Fatalln(err)
	}
}
