package main

import "net"
import "fmt"

func handleConn(conn net.Conn) {
	fmt.Println("host connected")
}

func main() {
	ln, err := net.Listen("tcp", ":10000")
	if err != nil {
		fmt.Printf("%v\n", err.Error())
	}

	for {
		conn, err := ln.Accept()
		if err != nil {
			fmt.Printf("%v\n", err.Error())
		}

		go handleConn(conn)
	}
}
