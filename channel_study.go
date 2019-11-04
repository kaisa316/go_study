package main

import (
	"fmt"
	"time"
)

//第一行<-c接受数据和最后一样接收数据完全不一样
//如果想要程序执行完毕，应该在最后一行<-c
//<-接收信息后，发送消息的goroutine会从block状态变为unblock状态；如果是main goroutine unblock后，调度器会开始调用它，并结束程序；所以如果放在第一行，那么只能for几十行；但如果for之后再接收消息(<-放在最后一行),那么greet就可以完整执行了
func Greet(c chan string) {
	//fmt.Println("hello " + <-c)
	for i := 0; i < 10000; i++ {
		fmt.Println(i)
	}
	fmt.Println("hello " + <-c)
}

//1.start main goroutine
//2.create channel
//3.start new goroutine,but main goroutine have control
//4.send msg to channel,main goroutine block,schedule other goroutine
func callGreet() {
	c := make(chan string)
	go Greet(c) //start new goroutine,ready state
	c <- "zhang sir"
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

func sleepFunc() {
	for i := 0; i < 10000; i++ {
		fmt.Println(i)
	}
}

//测试main goroutine在sleep 1ms后wake up，但正在执行的其他goroutine还没有执行完是否被会被main goroutine抢占。
//测试的结论是：main goroutine只要可用就会抢占
func callSleepFunc() {
	go sleepFunc()
	time.Sleep(time.Millisecond)
}

func sayHello(c chan string) {
	fmt.Println("say hello begin")
	<-c
	fmt.Println("say hello end")
	/*
		for i := 0; i < 10000; i++ {
			fmt.Println(i)
		}
	*/
}

func sayWorld(c chan string) {
	fmt.Println("say world begin")
	<-c
	fmt.Println("say world end")
}

func callSay() {
	fmt.Println("main goroutine start")
	c1 := make(chan string)
	c2 := make(chan string)
	go sayHello(c1)
	go sayWorld(c2)
	fmt.Println("send hello")
	c1 <- "hello"
	fmt.Println("send world")
	c2 <- "world"

	//time.Sleep(1 * time.Millisecond)
	fmt.Println("main end")
}

func main() {
	//main 函数本身就是一个main goroutine,而且当main goroutine执行完毕后，整个程序就会退出，不管有没有其他要执行的goroutine,如果没有阻塞block，那么goroutine会一直执行下去，即使有其他ready状态的goroutines也不会被轮转，他们会被活活饿死。
	//start main goroutine
	fmt.Println("main start")
	//callSayhi()
	//callBuffersize()
	//callSleepFunc()
	//callGreet()
	callSay()
	fmt.Println("main end")

}
