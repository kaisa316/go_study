package main

import (
	"fmt"
)

func sayHello(c chan string) {
	fmt.Println("say hello")
	<-c
}

func sayWorld(c chan string) {
	fmt.Println("say world")
	<-c
}

func main() {
	c1 := make(chan string)
	c2 := make(chan string)
	go sayHello(c1)
	go sayWorld(c2)
	select { //select block
	case c1 <- "hello":
		fmt.Println("send hello")
	case c2 <- "world":
		fmt.Println("send world")
	default:
		fmt.Println("default")
	}

}
