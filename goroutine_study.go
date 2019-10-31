/**
* go routine study
* go 的并发依靠go routine来实现，使用方式go xxx
* 多个goroutines之间依靠channel来进行通信
*
 */

package main

import (
	"fmt"
	"runtime"
	"time"
)

/**
* 并发执行一个函数
 */
func concurrent_study() {
	go say("world") // create a new goroutine
	say("hello")    // current goroutine
}

func say(s string) {
	for i := 0; i < 5; i++ {
		runtime.Gosched()
		fmt.Println(s)
	}
}

/**
* channel学习
* 多个不同的goroutine之间通过channel来进行通信
 */
func channel_study() {
	message := make(chan string)
	go func() { //new go routine
		message <- "hello ,我是goroutine num_1,收到了吗" //从这个goroutine中发送消息
	}()

	msg := <-message //current go routine,从这个goroutine中接受消息
	fmt.Println(msg) //goroutine num_2收到消息

}

//////10月28日，第二次学习///////
func hi(s string, ch chan string) {
	ch <- s
	fmt.Println(s)
}

func worker(done chan bool) {
	fmt.Println("working....")
	time.Sleep(2 * time.Second)
	fmt.Println("done")
	done <- true
}

func worker2(done chan bool) {
	fmt.Println("2 working")
	time.Sleep(2 * time.Second)
	fmt.Println("2 done")
	done <- true
}

func pln(s string, ch chan bool) {
	fmt.Println(s)
	time.Sleep(time.Second)
	fmt.Println("done")
	ch <- true
}

func main() {
	/*ch := make(chan string)
	for i := 0; i < 3; i++ {
		go hi("go world", ch)
	}
	<-ch
	fmt.Println("done ")*/
	/*done := make(chan bool)
	go worker(done)
	go worker2(done)
	<-done*/
	ch := make(chan bool)
	go pln("hello", ch)
	go pln("world", ch)
	go pln("nihao", ch)
	<-ch
}
