package main

import (
	"bufio"
	"fmt"
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

func handle(conn net.Conn) {
	read(conn)
}

func read(conn net.Conn) {
	scan := bufio.NewScanner(conn)
	for scan.Scan() {
		ln := scan.Text()
		if ln == ""{
			fmt.Println("end of the request")
			break
		}
		fmt.Println(ln)
	}
	defer conn.Close()
	fmt.Println("Code got here.")
	_, err := io.WriteString(conn, "I see you connected.")
	if err != nil {
		log.Fatalln(err)
	}
}
