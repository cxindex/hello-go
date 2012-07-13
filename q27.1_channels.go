package main

var c = make(chan int)

func lp() {
	for i := 0; i < 10; i++ {
		println(i)
	}
	c <- 1
}

func main() {
	go lp()
	println("Got from channel:", <-c)
}
