package main

import (
	"fmt"
	"net"
)

func main() {
	fmt.Println("client starting...")
	conn, err := net.Dial("tcp", "127.0.0.1:12345")
	defer conn.Close()
	if err != nil {
		fmt.Println(err)
	}
}
