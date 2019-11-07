package main

import (
	"fmt"
	"time"
)

var start time.Time

func init() {
	start = time.Now()
}

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

}

/////////rungo series ////////////
func service1(c chan string) {
	fmt.Println("service1 start", time.Since(start))
	for i := 0; i < 100; i++ {
		fmt.Println("i=", i)
	}
	time.Sleep(3 * time.Second)
	c <- "from service1 "
}

func service2(c chan string) {
	fmt.Println("service2 start", time.Since(start))
	for i := 0; i < 100; i++ {
		fmt.Println("K=", i)
	}
	time.Sleep(5 * time.Second)
	c <- "from service2 "
}

//阻塞式执行
func callService() {
	c1 := make(chan string)
	c2 := make(chan string)

	go service1(c1)
	go service2(c2)

	select {
	case res1 := <-c1:
		fmt.Println("response from service1 ", res1, time.Since(start))
	case res2 := <-c2:
		fmt.Println("response from service2 ", res2, time.Since(start))
	case <-time.After(2 * time.Second): //超时2秒执行这个
		fmt.Println("timeout ,no response receive")
	}

}

//非阻塞式执行
func callServiceUnblock() {
	c1 := make(chan string, 2) //buffer size 2
	c2 := make(chan string, 2) //buffer size 2

	c1 <- "value1"
	c1 <- "value2"
	c2 <- "value1"
	c2 <- "value2"

	select {
	case res1 := <-c1:
		fmt.Println("response from service1 ", res1, time.Since(start))
	case res2 := <-c2:
		fmt.Println("response from service2 ", res2, time.Since(start))
	}

}

func main() {
	fmt.Println("main goroutine start")
	callService()
	//callServiceUnblock()
	fmt.Println("main end")
}
