package main

import (
	"bytes"
	"fmt"
	"log"
	"net"
	//"net/http"
)

func main() {
	l, err := net.Listen("tcp", ":8080")
	if err != nil {
		panic(err)
	}
	fmt.Println("Listening on port 8080")

	for {
		conn, err := l.Accept()
		if err != nil {
			log.Fatal(err)
		}
		go handleConnection(conn)
	}
}

func handleConnection(conn net.Conn) {
	defer conn.Close()

	buf := bytes.NewBuffer(nil)
	n, err := buf.ReadFrom(conn)
	if err != nil {
		log.Println(err)
	}

	fmt.Printf("Read %d bytes: %s\n", n, buf.String())

	//fmt.Println(conn.LocalAddr(), conn.RemoteAddr())
	//conn.Write([]byte("Hello world!\n"))
}
