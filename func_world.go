package main

import (
	"fmt"
	"runtime"
)

/**
* 函数学习
* param
* return
* named return
 */
func funcStudy(name string, age int) (cityid int, cityname string) {
	defer deferStudy()
	fmt.Println(name, age)
	cityid = 1476
	cityname = "chifeng"
	return
}

/**
* defer：函数结束时要做的收尾工作
 */
func deferStudy() {
	fmt.Println("我来收尾了")
}

/**
* 函数做参数学习
 */
func paramFuncMain(age int, f func(name string)) {
	f("my name is yang")
	fmt.Printf("my age is %d", age)
	where()
}

/**
* 函数做参数学习
 */
func paramFuncStudy(name string) {
	fmt.Println(name)
}

func where() {
	_, file, line, _ := runtime.Caller(1)
	fmt.Println(file, line)
}
