//network test
package main

import (
	"bufio"
	"fmt"
	"io"
	"net"
	"os"
)

func main() {
	conn, err := net.Dial("tcp", "127.0.0.1:6666")
	defer conn.Close()
	if err != nil {
		fmt.Println("ERR:", err)
		return
	}
	// conn.Write([]byte("hello"))
	reader := bufio.NewReader(os.Stdin)
	for {
		s, rerr := reader.ReadString('\n')
		fmt.Fprintf(conn, "%s", s)
		if rerr == io.EOF {
			break
		}
	}
	fmt.Fprintf(conn, "END")

	fmt.Println(conn, err)
}
