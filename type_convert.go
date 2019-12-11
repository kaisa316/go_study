package main

import (
	"fmt"
	"strconv"
)

func main() {
	var i int
	i = 3
	s := string(i) //这样不行啊
	//fmt.Sprintf("%T,%v,xxxx", s, s)
	//fmt.Printf("%T,%v", i, i)
	fmt.Printf("%T,%v", s, s)
	fmt.Println()
	//用这个方式
	ss := strconv.Itoa(i)
	fmt.Printf("%T,%v", ss, ss)

	m := map[string]string{
		"3": "3333",
		"4": "4444",
	}
	r, ok := m[ss]
	if ok {
		fmt.Println("hello", r)
	}

	sss := fmt.Sprintf("%s,%d", i, i)
	fmt.Println(sss)
	//rr, oko := m[sss]
	//if oko {
	//fmt.Println("world", rr, oko)
	//} else {
	//fmt.Println("xxx")
	//}

	//fmt.Printf("%T,%v", i, i)
	//fmt.Println()
	//fmt.Println()
	//sss := fmt.Sprintf("%s", i)
	//fmt.Println(sss)
}
