package main

import (
	"fmt"
	//"time"
)

func sayHello(c chan string) {
	<-c
	/*
		for i := 0; i < 10000; i++ {
			fmt.Println(i)
		}
	*/
}

func sayWorld(c chan string) {
	//time.Sleep(time.Millisecond)
	<-c
}

func sayYou(c chan string) {
	c <- "see you"
}

func sayYou2(c chan string) {
	c <- "see you 2"
}

func blank() {
	fmt.Println("exec blank")
}

func callMe() {
	fmt.Println("main goroutine start")
	c1 := make(chan string)
	c2 := make(chan string)
	c3 := make(chan string)
	c4 := make(chan string)

	go sayYou2(c4)
	go sayHello(c1)
	go sayYou(c3)
	go sayWorld(c2)
	//go blank()
	//select只会执行其中一个case
	select { //select block
	case c1 <- "hello":
		fmt.Println("case:hello")
	case c2 <- "world":
		fmt.Println("case:world")
	case <-c3:
		fmt.Println("case:c3")
	case <-c4:
		fmt.Println("case:c4")
		//default: //if all the "case" 都没有准备好, execute default
		//time.Sleep(time.Millisecond)
		//fmt.Println("case:default")
	}
	//time.Sleep(1 * time.Millisecond)
	fmt.Println("main end")

}

func main() {
}
