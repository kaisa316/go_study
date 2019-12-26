package main

import (
	"fmt"
	"time"
)

func main() {
	//fmt.Println(time.Now())
	//fmt.Println(time.Parse("2006-01-02 15:04:05", time.Now()))
	//s, _ := time.ParseInLocation("2006-01-02 15:04:05", time.Now(), time.Local)
	formatDate()

}

//格式化当前时间为年月日时分秒Y-m-d H:i:s
func formatDate() {
	ss := time.Now().Format("2006-01-02 15:04:05")
	fmt.Println(ss)
}
