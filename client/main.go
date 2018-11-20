package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

const addr = "127.0.0.1:8080"

func main() {
	conn, err := net.Dial("tcp", addr)
	if err != nil {
		fmt.Printf("Failed to connect %s for reason: %s \n", addr, err.Error())
		os.Exit(1)
	}
	defer conn.Close()

	for {
		fmt.Print("> ")
		reader := bufio.NewReader(os.Stdin)
		t, err := reader.ReadString('\n')
		if err != nil {
			fmt.Printf("Failed to read from input: %s", err.Error())
			break
		}
		fmt.Printf("Started Connection with %s \n", addr)
		conn.Write([]byte(t))

		res := make([]byte, 1024)
		n, err := conn.Read(res)
		if err != nil {
			fmt.Printf("Failed to read data returned from server: %s", err.Error())
			break
		}
		fmt.Print(string(res[:n]))
	}
}
