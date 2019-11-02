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
		fmt.Println("buffer size func ,current i :", i)
		num := <-c
		fmt.Println("input is ", num)
	}
}

func callBuffersize() {
	c := make(chan int, 3)
	go buffersize(c)
	c <- 1
	c <- 2
	c <- 3
	c <- 4 //overflow ,才会block current goroutine

}

func main() {
	//main 函数本身就是一个main goroutine,而且当main goroutine执行完毕后，整个程序就会退出，不管有没有其他要执行的goroutine,如果没有阻塞block，那么goroutine会一直执行下去，即使有其他ready状态的goroutines也不会被轮转，他们会被活活饿死。
	//start main goroutine
	fmt.Println("main start")
	//callSayhi()
	callBuffersize()
	fmt.Println("main end")

}
