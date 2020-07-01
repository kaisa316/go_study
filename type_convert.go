package main

import (
	"fmt"
	"strconv"
)

func main() {
	//interfaceConvert()
	concreteConvert()
}

func otherMain() {
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
	//
	//
	//fmt.Println(sss)

}

//泛类型转换
func interfaceConvert() {
	var i interface{} = 7
	result := i.(int) //success
	fmt.Println(result)

	//str := i.(string) //fail,panic: interface conversion: interface {} is int, not string
	//fmt.Println(str)  //类型不同，转换失败

	str, ok := i.(string) //fail,但不会发生panic 错误
	fmt.Println(str, ok)
}

//具体类型转换
func concreteConvert() {
	i := 7
	s := string(i) //success
	output := fmt.Sprintf("%T,%v", s, s)
	fmt.Println(output)

	s2 := "7"
	//i2 := int(s2) //fail,cannot convert sk (type string) to type int
	//output2 := fmt.Sprintf("%T,%v", i2, i2)
	//fmt.Println(output2)

	i3, ok := strconv.Atoi(s2)
	output3 := fmt.Sprintf("%T,%v", i3, i3)
	fmt.Println(output3, "error:", ok)
}
