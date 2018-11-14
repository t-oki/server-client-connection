package main

import (
	"fmt"
	"log"
	"net"
)

func main() {
	listener, err := net.Listen("tcp", "127.0.0.1:8080")
	if err != nil {
		log.Fatal(err)
	}
	defer listener.Close()
	log.Println("Listening on 127.0.0.1:8080..")

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Fatal(fmt.Sprintf("Connection Error: %s", err))
		}
		defer conn.Close()

		go func() {
			remoteAddr := conn.RemoteAddr().String()
			log.Println(fmt.Sprintf("Client connected from %s", remoteAddr))

			buf := make([]byte, 1024)
			n, err := conn.Read(buf)
			if err != nil {
				log.Fatal("Failed to read from connection")
			}
			log.Println(fmt.Sprintf("%s", buf[:n]))

			res := fmt.Sprintf("Hello, %s", remoteAddr)
			conn.Write([]byte(res))
		}()
	}
}
