/**
* interface 接口学习
 */

package main

import (
	"fmt"
)

type Animal interface {
	head(desc string) (color string, desc_rtn string)
	body(desc string) (color string, desc_rtn string)
}

type horse struct {
	color string
}

func (ma horse) head(desc string) (color string, desc_rtn string) {
	color = ma.color
	desc_rtn = desc
	return
}

func (ma horse) body(desc string) (color string, desc_rtn string) {
	color = ma.color
	desc_rtn = desc
	return
}

func test(t Animal) {
	fmt.Println(t.head("this is head"))
	fmt.Println(t.body("this is body"))
}

func interfaceStudy() {
	black := new(horse)
	black.color = "black horse"
	test(black)
}

type shape interface {
	area()
}

type Rectangle struct {
	width  int
	height int
	note   string
}

type Circle struct {
	radius int
	note   string
}

func (r Rectangle) area() {
	fmt.Println("i am Rectangle area")
}

func (c Circle) area() {
	fmt.Println("i am Circle area")
}

func showArea(s shape) {
	s.area()
}

//以国家为例
type country struct {
	name string  //国名
	flag string  //国旗
	area float64 //国土面积
}

type china struct {
	country
	yellowRiver string //黄河
}

type american struct {
	country
	amzRiver string //亚马逊河
}

type show interface {
	showRiver()
	showFlag()
}

func (c china) showRiver() {
	fmt.Println(c.yellowRiver)
}

func (c china) showFlag() {
	fmt.Println(c.flag)
}

func (a american) showRiver() {
	fmt.Println(a.amzRiver)
}

func (a american) showFlag() {
	fmt.Println(a.flag)
}

func showCountry(s show) {
	s.showRiver()
	s.showFlag()
}

func main() {
	c := china{}
	c.name = "china"
	c.flag = "五星红旗"
	c.yellowRiver = "黄河"

	a := american{}
	a.name = "美国"
	a.flag = "星条旗"
	a.amzRiver = "亚马逊河"

	showCountry(c)
	showCountry(a)
}
