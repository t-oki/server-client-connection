package main

import (
	"fmt"
	"log"
	"net"
)

func main() {
	conn, err := net.Dial("tcp", "127.0.0.1:8080")
	if err != nil {
		log.Fatal("Failed to connect to 127.0.0.1:8080")
	}
	defer conn.Close()

	msg := fmt.Sprintf("Hello, %s", conn.RemoteAddr())
	conn.Write([]byte(msg))

	res := make([]byte, 1024)
	n, _ := conn.Read(res)
	log.Println(string(res[:n]))
}
