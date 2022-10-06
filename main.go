package main

import (
	"bufio"
	"fmt"
	"io"
	"net"
	"os"
)

var port = ":5000"

func echo(conn net.Conn) {
	r := bufio.NewReader(conn)
	for {
		line, err := r.ReadBytes(byte('\n'))
		switch err {
		case nil:
			break
		case io.EOF:
		default:
			fmt.Println("ERROR", err)
		}
		conn.Write(line)
	}
}

func main() {
        fmt.Println("Starting server...")

	l, err := net.Listen("tcp", port)
        fmt.Println("Waiting for players...")

	if err != nil {
		fmt.Println("ERROR", err)
		os.Exit(1)
	}

	for {
		conn, err := l.Accept()
		if err != nil {
			fmt.Println("ERROR", err)
			continue
		}
            // Handle connections in a new goroutine.
            // go myHandler(conn)
            go echo(conn)
	}

}