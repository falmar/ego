package main

import (
	"net"
)

func main() {

	ln, err := net.Listen("tcp",":8081")

	if err != nil {
		panic(err)
	}
	defer ln.Close()

	for {
		conn, err := ln.Accept()

		if err != nil {
			panic(err)
		}

		conn.Write([]byte("What's Up\n"))

		conn.Close()
	}

}
