package main

import "fmt"

func lp(c, q chan int) {
	for {
		select {
		case get := <-c:
			fmt.Println(get)
		case <-q:
			println("Q")
			break
		}
	}
}

func main() {
	c := make(chan int)
	q := make(chan int)
	go lp(c, q)
	for i := 0; i < 10; i++ {
		c <- i
	}
	q <- 0
}
