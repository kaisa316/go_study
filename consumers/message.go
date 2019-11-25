package consumers

import (
	"fmt"
	"time"
)

//10万个任务，同时消费N个任务
func Msg() {
	start := time.Now()
	total := 100000 //总次数
	multiNum := 1   //并发数量
	chanBuffer := make(chan int, multiNum)
	for i := 0; i < total; i++ {
		chanBuffer <- i
		go work(i, chanBuffer)
	}
	fmt.Println(time.Since(start))
}

func work(index int, ch chan int) {
	//do worker
	fmt.Println(index)
	<-ch
}
