package main

import "fmt"

func lp(c chan int) {
	for {
		get := <-c
		fmt.Println(get)
	}
}

func main() {
	c := make(chan int)
	go lp(c)
	for i := 0; i < 10; i++ {
		c <- i
	}

}
