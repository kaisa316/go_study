/**
* map 学习
 */
package main

import (
	"fmt"
)

func main() {
	//第一种方式
	var test map[string]string
	test = make(map[string]string) //初始化
	test["name"] = "zhangsan"
	test["age"] = "23"
	fmt.Println(test)

	//第二种方式
	mapB := map[string]string{}
	mapB["name"] = "lisi"
	mapB["city"] = "chifeng"
	fmt.Println(mapB)
	//第三种方式
	mapC := make(map[string]string)
	mapC["name"] = "wangwu"
	fmt.Println(mapC)

}
