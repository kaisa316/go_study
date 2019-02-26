/*
* go study,基本类型的学习
* 通用原则：
* 1. 无论是变量，方法，常量都采用驼峰式写法，而不是_下划线方式分割
 */
package main

import (
	"fmt"
	//"os"
	//"runtime"
	"strings"
	"time"
)

var (
	//全局变量允许声明但不使用，局部变量，声明之后必须使用
	personName   string = "super man"
	personAge    int    = 32
	personSchool string
)

const (
	piVal float32 = 3.1415
	hours int     = 24
)

func main() {
	//studyVar()
	//name, _ := GetName("zhangsan")
	//fmt.Println(name)
	//operator()
	//typeConvert()
	//studyConst()
	//charStudy()
	//dateStudy()
	//pointerStudy()
	//ifStudy()
	//forStudy()
	//switchStudy()
	//fmt.Println(funcStudy("zhangsan", 23))
	//paramFuncMain(99, paramFuncStudy)
	//arrayStudy()
	//sliceStudy()
	//mapStudy()
	//structStudy()
	interfaceStudy()
}

/*
* 变量的学习
* 两种方式
* 第一种，var形式
* 第二种，:= 形式
 */
func studyVar() {
	//局部变量，定义了会报错，全局变量定义了不用，也不会报错
	//var name string            //decalre
	//var address string = "内蒙古" //declare && init value
	var hometown = "赤峰市" //omit type
	age := 23            // := style,局部变量首选方式
	//var ff = name + address + hometown //string combine
	yangchang := "进驻"
	yangchang, hometown = hometown, yangchang //exchange variable value
	fmt.Println("yangchang:"+yangchang, "hometown"+hometown, &age, personSchool)
	//path := os.Getenv("PATH")
	//fmt.Println(path)
	//runtime := runtime.GOOS
	//s := fmt.Sprintf("hello ,%s", runtime)
	//fmt.Println(s)
	//cc := 99
	//fmt.Println(cc)
}

/**
* 常量学习
 */
func studyConst() {
	const PI = 3.14
	fmt.Println(piVal)
}

/**
* 函数的学习
*
* 参数:名称，类型
* 返回值 名称，类型 ，可以多个返回值
 */
func GetName(name string) (string, int) {
	return name, 23
}

/**
* 类型转换
 */
func typeConvert() {
	var height = 8848.66
	var f = int(height)
	fmt.Println(f)
}

func operator() {
	i := 0
	i++
	j := i + 5
	fmt.Println(i, j)
}

/**
* 字符类型学习
 */
func charStudy() {
	//var c int = '张'
	//c := '王'
	c := 'a'
	fmt.Println("char c is", c)
	fmt.Printf("char c is%U", c)
}

/**
* 字符串类型学习
 */
func stringStudy() {
	//sfirst := "第一种形式,解释型"
	/*ssecond := `第二种方式
	raw string`
	sthird := "abc"
	fmt.Println(sfirst)
	fmt.Println(ssecond)
	fmt.Println(sthird[1])*/
	ss := "first blood"
	slice := strings.Split(ss, " ")
	fmt.Println(strings.Join(slice, "--"))
}

/**
* 时间、日期学习
 */
func dateStudy() {
	fmt.Println(time.Now())
	fmt.Println(time.Now().Format("20060102"))
}

/**
* 指针学习
* 内存就像快递柜，每个柜子当中放有快递包裹（value），和他的坐标[3,2](内存地址)
 */
func pointerStudy() {
	a := 5
	s := "zhangsan"
	b := &a
	ptr := &s
	var m *string
	m = &s
	fmt.Println(b, ptr, m)

}
