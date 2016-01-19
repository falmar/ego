package main

import (
	"net"
	"bufio"
	"strings"
	"sync"
	"fmt"
)

func main() {

	ln, err := net.Listen("tcp", ":8082")

	if err != nil {
		panic(err)
	}
	defer ln.Close()

	for {
		conn, err := ln.Accept()

		if err != nil {
			panic(err)
		}

		go handle(conn)
	}
}

var redis map[string]string = make(map[string]string)
var mutx sync.Mutex

func handle(conn net.Conn) {
	defer conn.Close()
	scanner := bufio.NewScanner(conn)
	Lopper:
	for scanner.Scan() {
		cmds := strings.Split(scanner.Text(), " ")
		switch strings.ToLower(cmds[0]) {
		case "get":
			rGet(conn, cmds[1])
		case "set":
			rSet(conn, cmds[1], cmds[2])
		case "del":
			rDel(conn, cmds[1])
		case "quit":
			break Lopper
		default:
			conn.Write([]byte(fmt.Sprintf("Invalid command: %s ", cmds[0]) + "\n"))
		}
	}
}
func rGet(conn net.Conn, Key string) {
	mutx.Lock()
	defer mutx.Unlock()
	val, ok := redis[Key]
	if ok {
		conn.Write([]byte(val + "\n"))
	} else {
		conn.Write([]byte("-1\n"))
	}
}

func rSet(conn net.Conn, Key, Value string) {
	mutx.Lock()
	defer mutx.Unlock()
	redis[Key] = Value
	conn.Write([]byte("1\n"))
}

func rDel(conn net.Conn, Key string) {
	mutx.Lock()
	defer mutx.Unlock()
	_, ok := redis[Key]
	if ok {
		delete(redis, Key)
		conn.Write([]byte("1\n"))
	} else {
		conn.Write([]byte("-2\n"))
	}
}
