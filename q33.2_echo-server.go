//here we can read and send something back. concurrent
//based on q33.1_echo-server_v2.go
package main

import (
	"fmt"
	"net"
	"bufio"
	"os"
	"io/ioutil"
)

func ConnHandler(conn *net.Conn, data *[]byte) {
	fmt.Println("Start conn:", conn)
	defer (*conn).Close()
	(*conn).Write([]byte(*data))

	q := make(chan int)
	b := make([]byte, 1024)
	for {
		go func() {
			n, _ := (*conn).Read(b)
			if n == 0 {
				fmt.Println("End conn:", conn)
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
	(*conn).Close()
	return
}

func main() {
	l, err := net.Listen("tcp", ":6666")
	if err != nil {
		fmt.Println(err)
		return
	}
	r := bufio.NewReader(os.Stdin)
	b, _ := ioutil.ReadAll(r)
	fmt.Println(string(b))
	for {
		conn, err := l.Accept()
		if err != nil {
			fmt.Println(err, conn)
			return
		}
		go ConnHandler(&conn, &b)
	}
}
