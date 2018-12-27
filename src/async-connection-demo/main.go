package main

import (
	"log"
	"net"
	"os"
	"time"
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
	defer conn.Close()
	for {
		//Setting a timeout for the connection here
		conn.SetReadDeadline(time.Now().Add(5 * time.Second))

		var buf [128]byte

		n, err := conn.Read(buf[:])

		if err != nil {
			log.Printf("Finished with err = %v", err)
			return
		}

		os.Stderr.Write(buf[:n])
	}
}
