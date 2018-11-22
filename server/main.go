package main

import (
	"log"
	"net"
)

const addr = "127.0.0.1:8080"
const udpaddr = "127.0.0.1:8081"

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

		conn.Write([]byte("tcp"))
	}
}

func listenTCP() {
	listener, err := net.Listen("tcp", addr)
	if err != nil {
		log.Fatal(err)
	}
	defer listener.Close()
	log.Printf("Listening TCP on %s...", addr)

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

func listenUDP() {
	conn, err := net.ListenPacket("udp", udpaddr)
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()
	log.Printf("Connected UDP on %s...", addr)

	buf := make([]byte, 1024)
	for {
		n, remoteAddr, err := conn.ReadFrom(buf)
		if err != nil {
			log.Fatal(err)
		}
		log.Printf("%s: %s", remoteAddr, buf[:n])

		go func() {
			conn.WriteTo([]byte("udp"), remoteAddr)
		}()
	}
}

func main() {
	go listenTCP()
	listenUDP()
}
