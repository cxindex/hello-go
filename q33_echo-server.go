package main

import (
	"fmt"
	"net"
)

func ConnHandler(conn *net.Conn) []byte {
	b := make([]byte, 1024)
	for {
		n, _ := (*conn).Read(b)
		if n == 0 {
			break
		}
		fmt.Printf("!%s", string(b))
	}
	(*conn).Close()
	return b
}

func main() {
	l, _ := net.Listen("tcp", ":6666")
	conn, err := l.Accept()
	if err != nil {
		fmt.Println(err)
		return
	}
	ConnHandler(&conn)
}
