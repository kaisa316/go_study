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
* 接收者是指针pointer
 */
func (r *Rectangle) setLong(v int) {
	r.long = v
}

func (r Rectangle) setWidth(w int) {
	r.width = w
	fmt.Println(r.width)
}

func main() {
	//rec := Rectangle{height: 5, width: 6}
	var rec Rectangle
	rec.height = 5
	rec.width = 5
	//areaRec := rec.area()
	//fmt.Println(areaRec)

	rec2 := Rectangle{}
	rec2.height = 5
	rec2.width = 8

	rec2.setLong(4)
	rec2.setWidth(2)
	fmt.Println(rec2.width)
	fmt.Println(rec2.long)
}
