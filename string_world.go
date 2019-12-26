package main

import (
	"fmt"
)

func main() {
	//s := "世界"
	//ss := firstChar(s)
	//fmt.Println(ss)

	isEmpty()
}

//获取字符串首字符，比如世界，获取世界
func firstChar(str string) string {
	return string([]rune(str)[0])
}

//域名
func domain() string {
	url := "http://api%s.baidu.com"
	suffix := ""
	if true { //测试环境
		suffix = ".ceshi"
	}
	url = fmt.Sprintf(url, suffix)
	return url

}

func isEmpty() {
	s := "0"
	fmt.Println(s == "0")
	if s != "0" {
		fmt.Println(222)
		fmt.Println("s !='0'", s)
	}
}
