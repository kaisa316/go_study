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
