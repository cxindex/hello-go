//here we can read and write. concurrent
package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

func ConnHandler(conn *net.Conn) []byte {
	println("START")
	defer (*conn).Close()
	reader := bufio.NewReader(os.Stdin)
	q := make(chan int)
	b := make([]byte, 1024)
	for {
		go func() {
			s, _ := reader.ReadString('\n')
			(*conn).Write([]byte(s))
			q <- 1
		}()
		go func() {
			n, _ := (*conn).Read(b)
			if n == 0 {
				println("BREAK")
				q <- -1
				return
			}
			fmt.Printf("%s", string(b[:n]))
			q <- 1
		}()
		if <-q == -1 {
			break
		}
	}
	return b
}

func main() {
	l, err := net.Listen("tcp", ":6666")
	if err != nil {
		fmt.Println(err)
	}
	conn, err := l.Accept()
	if err != nil {
		fmt.Println(err)
		return
	}
	ConnHandler(&conn)
}
