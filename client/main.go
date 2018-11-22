package main

import (
	"fmt"
	"log"
	"net"
	"os"
)

const addr = "127.0.0.1:8080"
const udpaddr = "127.0.0.1:8081"

func dialTCP() {
	conn, err := net.Dial("tcp", addr)
	if err != nil {
		fmt.Printf("Failed to connect %s for reason: %s \n", addr, err.Error())
		os.Exit(1)
	}
	defer conn.Close()
	fmt.Printf("Started TCP Connection with %s \n", addr)

	go func() {
		for i := 0; i < 30; i++ {
			conn.Write([]byte(fmt.Sprintf("TCP, %d 度目", i)))
		}
	}()

	for {
		res := make([]byte, 4096)
		n, err := conn.Read(res)
		if err != nil && err.Error() == "EOF" {
			log.Println("disconnected")
			break
		}
		if err != nil {
			fmt.Printf("Failed to read data returned from server: %s", err.Error())
		}
		fmt.Println(string(res[:n]))
	}
}

func dialUDP() {
	conn, err := net.Dial("udp", udpaddr)
	if err != nil {
		fmt.Printf("Failed to connect %s for reason: %s \n", udpaddr, err.Error())
		os.Exit(1)
	}
	defer conn.Close()
	fmt.Printf("Started UDP Connection with %s \n", udpaddr)

	go func() {
		for i := 0; i < 30; i++ {
			conn.Write([]byte(fmt.Sprintf("UDP, %d 度目", i)))
		}
	}()

	for {
		res := make([]byte, 1024)
		n, err := conn.Read(res)
		if err != nil {
			fmt.Printf("Failed to read data returned from server: %s", err.Error())
		}
		fmt.Println(string(res[:n]))
	}
}

func main() {
	go dialTCP()
	dialUDP()
}
