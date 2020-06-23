package main

import "fmt"

func output(c chan int) {
	for i:=0;i<10;i++ {
		fmt.Println(i)
		c <- i
	}
}

func main() {
	c := make(chan int)
	fmt.Println("main start")
	go output(c)
	for {
		val,ok:= <-c 
		if ok {
			fmt.Println(val,ok)
		} else {
			fmt.Println("channel close")
			break
		}
		
	}
	fmt.Println("main end")
}