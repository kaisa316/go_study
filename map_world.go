/**
* map 学习
 */
package main

import (
	"fmt"
)

func main() {
	//mapSlice()
	mapSlice3()
}

func practice() {
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

func mapSlice() {
	//["key"][0]=1
	//["key"][1]=10
	//["key"][2]=100
	var m map[string][]int
	m = make(map[string][]int)
	var s []int
	arr := []int{1, 10, 100, 1000, 99, 88}
	for _, val := range arr {
		s = append(s, val) //append itself , 这是关键
		m["id"] = s
	}
	fmt.Println(m)

}
func sl() {
	m := map[string]map[int][]int{}
	//m["a"][1] = append(m["a"][1], 55)
	fmt.Println(m)

}
func mapSlice2() {
	// db 返回的形式:
	//["key"][2][]=10
	//["key"][2][]=100
	//["key"][2][]=1000
	//["key2"][1][]=2
	//["key2"][1][]=22
	//["key2"][1][]=222
	//最终的想要的结果形式:
	//["key"][2] = [10,100,1000]
	//["key2"][1] = [2,22,222]
	var name map[string]map[int][]int //map map int_slice
	name = make(map[string]map[int][]int)
	data := []int{5, 10, 15, 20}
	secData := make(map[int][]int)
	secData[2] = data
	name["zhangsan"] = secData
	fmt.Println(name)

}

func mapSlice3() {
	dbData := []map[string]string{
		{"score": "1", "name": "zhangsan", "kecheng": "yuwen"},
		{"score": "10", "name": "zhangsan", "kecheng": "yuwen"},
		{"score": "100", "name": "zhangsan", "kecheng": "math"},
		{"score": "1000", "name": "zhangsan", "kecheng": "math"},
		{"score": "2", "name": "lisi", "kecheng": "yuwen"},
		{"score": "22", "name": "lisi", "kecheng": "yuwen"},
		{"score": "222", "name": "lisi", "kecheng": "math"},
		{"score": "2222", "name": "lisi", "kecheng": "math"},
	}
	//["zhangsan"]["yuwen"] = ["1","10"]
	//["zhangsan"]["math"] = ["100","1000"]
	//map map slice
	//这种形式，第二个map也不要实例化才行,否则报错
	//姓名，课程，分数
	var result map[string]map[string][]string
	result = make(map[string]map[string][]string) //map 第一种初始化形式
	// kecheng := map[string][]string{}              //第二种map 初始化形式
	// score := []string{}
	for _, val := range dbData {
		name := val["name"]
		kechengVal := val["kecheng"]
		score := result[name][kechengVal]
		score = append(score, val["score"])
		kecheng := result[name]
		if kecheng == nil {
			kecheng = make(map[string][]string)
		}
		// result[name][kechengVal] = append(result[name][kechengVal], val["score"])
		//这样会报错,内部map没有初始化
		kecheng[kechengVal] = score
		result[name] = kecheng
	}

	fmt.Println(result)
	//

}
