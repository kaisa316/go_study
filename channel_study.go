package main

import (
	"fmt"
)

func Greet(c chan string) {
	fmt.Println("world")
	fmt.Println("hello " + <-c)
}

func circle(c chan int) {
	for i := 0; i < 10; i++ {
		fmt.Println("i am circle")
		c <- i
	}
	close(c)
}

func callCircle() {
	c := make(chan int)
	go circle(c) //start goroutine
	//block/unblock of main goroutine util channel close
	for {
		val, ok := <-c
		if ok == false {
			fmt.Println(val, ok, "loop break")
			break //exit break loop
		} else {
			fmt.Println(val, ok)
		}
	}
	/*
		for ok := range c {
			fmt.Println(ok)
		}
	*/
}

func sayHi(c chan string) {
	<-c //get john
	fmt.Println("first hello")
	<-c //channel中什么都没有，block current goroutine
	fmt.Println("second hello")
}

func callSayhi() {
	c := make(chan string)
	go sayHi(c)
	c <- "john"
}

func buffersize(c chan int) {
	for i := 0; i < 10; i++ {
		//fmt.Println("the length is  ", len(c))
		fmt.Println("xxx")
		num := <-c
		fmt.Println("input is ", num)
	}
}

func callBuffersize() {
	c := make(chan int, 3)
	go squares(c)
	c <- 1
	c <- 2
	c <- 3
	c <- 4 //overflow ,才会block current goroutine

}

func main() {
	fmt.Println("main start")
	//callSayhi()
	fmt.Println("main end")

}
