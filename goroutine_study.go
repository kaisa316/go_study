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
)

func goroutine_study() {
	channel_study()
}

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
	fmt.Println(msg)

}
