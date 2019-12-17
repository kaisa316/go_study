package main

import (
	"fmt"
)

func main() {
	s := "世界"
	ss := firstChar(s)
	fmt.Println(ss)

}

//获取字符串首字符，比如世界，获取世界
func firstChar(str string) string {
	return string([]rune(str)[0])
}
