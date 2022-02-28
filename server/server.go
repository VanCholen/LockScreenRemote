package main

import (
	"fmt"
	"net"
	"os/exec"
)

func handleConnection(conn net.Conn) {
	defer conn.Close()
	cmd1 := exec.Command("cmd", "/c", `rundll32.exe user32.dll, LockWorkStation`)
	cmd1.Run()
}
func main() {
	fmt.Println("server starting...")
	server, err := net.Listen("tcp", "127.0.0.1:12345")
	if err != nil {
		fmt.Println(err)
	}
	for {
		conn, err := server.Accept()
		if err != nil {
			// handle error
			continue
		}
		go handleConnection(conn)
	}
}
