package main

import (
	"fmt"
)

type shape struct {
	width  int
	height int
}

/**
* 定义长方形类型
 */
type Rectangle struct {
	shape //匿名字段,继承
	long  int
}

/**
* 长方形的面积
 */
func (r Rectangle) area() (result int) {
	result = r.height * r.width
	return
}

func (r Rectangle) area2(width int, height int) (result int) {
	result = width * height
	return
}

func (r Rectangle) getVolume() (v int) {
	v = r.height * r.width * r.long
	return
}

/**
* 接收者是指针pointer ,传递的是对象的内存地址copy,接收者自身可以被改变
 */
func (r *Rectangle) setLong(v int) {
	r.long = v

}

/**
* 接收者是value ,传递的是对象的copy,只会改变copy自身
 */
func (r Rectangle) setWidth(w int) {
	r.width = w
}

func main() {
	rec2 := Rectangle{}
	rec2.height = 5
	rec2.width = 5

	rec2.setLong(14)
	fmt.Println(rec2.long) //output 14，指针改变了原始值

	rec2.setWidth(10)
	fmt.Println(rec2.width) //output 5,value 传值只会改变副本

	r := rec2.area()
	fmt.Println(r)
}
