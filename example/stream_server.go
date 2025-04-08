package main

import (
	"fmt"
	"io"
	"log"
	"net"
	"os"
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
	out, err := os.OpenFile("temp.txt", os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		panic(err)
	}

	var total int64

	for {
		n, err := io.CopyN(out, conn, 1024)
		total += n
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Println(err)
			break
		}
	}

	fmt.Printf("Read %d bytes: %s\n", total, out.Name())
}
