package main

import (
	"io"
	"log"
	"net"
	"os"
)

func main() {
	listener, err := net.Listen("tcp", "localhost:4242")

	if err != nil {
		log.Fatal(err)
	}

	//Keeps accepting new connection here
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Fatal(err)
		}

		go copyToStderr(conn)
	}
}

func copyToStderr(conn net.Conn) {
	n, err := io.Copy(os.Stderr, conn)
	log.Printf("Copied %d bytes; finished with err = %v", n, err)
}
