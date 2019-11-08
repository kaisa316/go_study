package main

import (
	"fmt"
)

type Person struct {
	name string
	age  int
}

func cal(num ...int) (sum int) {
	for _, val := range num {
		//fmt.Println(val)
		sum += val
	}
	return
}

type typeC struct {
}

///////
type typeCon interface {
	typeConvert()
}

func (t typeC) typeConvert() {
	i := 32
	j := float32(i)
	fmt.Printf("%T,%v\n", i, i)
	fmt.Printf("%T,%v", j, j)
}

func main() {
	p := Person{}
	p.age = 20
	p.name = "zhangsan"
	fmt.Println(p)

	sum := cal(1, 2, 3)
	fmt.Println(sum)

	//typeConvert()
	t := typeC{}
	t.typeConvert()
}
