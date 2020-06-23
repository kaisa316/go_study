package main

import (
	"fmt"
	"runtime"
	"sync"
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
	for i := 0; i < 3; i++ {
		fmt.Println("circle goroutine:i am circle ", i)
		c <- i
	}
	fmt.Println("before close")
	close(c)
	fmt.Println("after close")
}

func callCircle() {
	c := make(chan int)
	go circle(c) //start goroutine
	//block/unblock of main goroutine util channel close
	for {
		val, ok := <-c
		fmt.Println("main here")
		if ok == false {
			fmt.Println(val, ok, "loop break")
			break //exit break loop
		} else {
			fmt.Println("main goroutine:", val, ok)
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

func sender(c chan int) {
	c <- 1 //len 1,cap 3
	c <- 2 //len 2,cap 3
	c <- 3 //len 3,cap 3
	c <- 4 //goroutine block here
	close(c)
}

func callSender() {
	c := make(chan int, 3)
	go sender(c)
	for val := range c { //read values from c,block here
		fmt.Println("val is ", val)
	}
}

func senderUnclose(c chan int) {
	for i := 0; i < 4; i++ {
		fmt.Println(<-c)
	}
}

func callSenderUnclose() {
	c := make(chan int, 3)
	go senderUnclose(c)
	fmt.Println("active goroutines ", runtime.NumGoroutine())
	c <- 1
	c <- 2
	c <- 3
	c <- 4
	fmt.Println("active goroutines ", runtime.NumGoroutine())
	go senderUnclose(c)
	fmt.Println("active goroutines ", runtime.NumGoroutine())
	c <- 5
	c <- 6
	c <- 7
	c <- 8 //block here
	fmt.Println("active goroutines ", runtime.NumGoroutine())

}

//平方计算
func square(c chan int) {
	fmt.Println("square start")
	num := <-c
	fmt.Println("square end", num)
	//c <- num * num
}

//立方计算
func cube(c chan int) {
	fmt.Println("cube start")
	num := <-c
	fmt.Println("cube end", num)
	//c <- num * num * num
}

func callCaculate() {
	c1 := make(chan int)
	c2 := make(chan int)
	go square(c1)
	go cube(c2)
	fmt.Println("send to c1")
	c1 <- 3
	fmt.Println("send to c2")
	c2 <- 4
	fmt.Println("active goroutines:", runtime.NumGoroutine())
}

func waitLearn(wg *sync.WaitGroup, input int) {
	fmt.Println("execute waitLearn = ", input)
	wg.Done()
}

//select 是cases中的任意一个通了后就unblock
//waitgroup见名知意，是一组。必须所有case都准备好之后才能unblock
func callWaitLearn() {
	var wg sync.WaitGroup //create waitgroup (empty struct)
	for i := 0; i < 3; i++ {
		wg.Add(1) //increment counter
		go waitLearn(&wg, i)
	}
	wg.Wait() //block here
}

func chanClose() {
	c := make(chan int)
	go chanCloseBranch(c)
	c <- 3
	close(c)
	fmt.Println("end")
}

func chanCloseBranch(c chan int) {
	i, _ := <-c
	fmt.Println(i)
	c <- 5
	fmt.Println("kkk")
}

func main() {
	//routine,如果没有阻塞block，那么goroutine会一直执行下去，即使有其他ready状态的goroutines也不会被轮转，他们会被活活饿死。
	fmt.Println("main start")
	//callSayhi()
	//callBuffersize()
	//callSleepFunc()
	//callGreet()
	//callSay()
	//callSender()
	//callSenderUnclose()
	//callCaculate()
	//callWaitLearn()
	//fmt.Println("main end")
	callCircle()
	//chanClose()

}
