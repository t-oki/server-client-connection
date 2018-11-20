package main

import (
	"fmt"
	"log"
	"net"
)

const addr = "127.0.0.1:8080"

func main() {
	listener, err := net.Listen("tcp", addr)
	if err != nil {
		log.Fatal(err)
	}
	defer listener.Close()
	log.Printf("Listening on %s...", addr)

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Printf("Connection Error: %s", err)
			continue
		}
		defer conn.Close()

		go handleConnection(conn)
	}
}

func handleConnection(conn net.Conn) {
	remoteAddr := conn.RemoteAddr().String()
	log.Printf("Client connected from %s", remoteAddr)

	for {
		buf := make([]byte, 1024)
		n, err := conn.Read(buf)
		if err != nil && err.Error() == "EOF" {
			log.Printf("Client: %s disconnected", remoteAddr)
			break
		} else if err != nil {
			log.Println(err)
			log.Fatal("Failed to read from connection")
		}
		log.Printf("%s: %s", remoteAddr, buf[:n])

		res := fmt.Sprintf("Your word:  %s", buf[:n])
		conn.Write([]byte(res))
	}
}
