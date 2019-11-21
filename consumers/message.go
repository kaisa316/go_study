package consumers

import (
	"fmt"
)

func Msg() {
	total := 100000 //总次数
	multiNum := 10  //并发数量
	chanBuffer := make(chan int, multiNum)
	for i := 0; i < total/multiNum; i++ {
		go work(i, chanBuffer)
		chanBuffer <- i
	}
}

func work(index int, ch chan int) {
	//do worker
	fmt.Println(index)
	<-ch
}
