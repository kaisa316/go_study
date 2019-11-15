package main

import (
	"fmt"
	"github.com/kaisa316/nummanip/calc"
	cc "github.com/kaisa316/nummanip/v2/calc"
	yy "github.com/kaisa316/nummanip/v3/calc"
)

func main() {
	fmt.Println("hello")
	s := calc.Sum(1, 2, 3)
	fmt.Println(s)
	c := cc.Mul(3, 5)
	fmt.Println(c)
	yy.Mul(1, 2)
}
