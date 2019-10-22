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

/**
* base function
 */
func addFuc(a int, b int) int {
	return a + b
}

/**
* named return
* 命名返回
 */
func addFuc2(a, b int) (sum int) {
	sum = a + b
	return
}

/**
* named function
* 命名返回，多返回值
 */
func addFuc3(a int, b int) (sum int, cj int) {
	sum = a + b
	cj = a * b
	return
}

/**
* 可变参数
*
 */
func addFuc4(a ...int) (result int) {
	for key, val := range a {
		fmt.Println(key, val)
	}

	return 5
}

/**
* 传指针
 */
func addFuc5(a *int, b *int) {
	fmt.Println(*a)
	fmt.Println(&a)
	*a = 6 //dereference,更改地址所对应的content
	*b = 8
}

func main() {
	a := 3
	b := 4
	addFuc5(&a, &b) //传地址，a,b变量所对应的
	fmt.Println(a, b)

}
