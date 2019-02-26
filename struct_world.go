/**
* struct and method study
 */
package main

import (
	"fmt"
	"go_study/world"
)

//自定义类型 两种方式： 类型别名alias type and 结构struct
type yang struct {
	name string
	age  int
}

type intAlias int

type twoYang struct {
	yang //匿名字段相当于继承
	name string
	city string
}

/**
* struct 入口文件
 */
func structStudy() {
	//第一种，用new创建yang 类型
	/*yy := new(yang)
	yy.name = "name:zhangsansan"
	yy.age = 23
	yy.echoInfo()
	fmt.Println(yy)*/

	//第二种，类型初始化
	/*yy := yang{name: "wangwu", age: 5}
	yy.echoInfo()
	fmt.Println(yy)*/
	//基本类型alias
	/*m := new(intAlias)
	m.testInt()*/
	/*y := new(yang)
	y.name = "zhaoliu"
	y.age = 22
	testYang(y)*/

	/*yy := yang{name: "zhaoliu val", age: 22}
	testYangVal(yy)
	*/
	//factoryTest()
	extendTest()
}

/**
* method
 */
func (yy *yang) echoInfo() {
	fmt.Println(yy.name, yy.age)
	yy.name = "name are changed"
}

/**
* 基本类型alias
* method
 */
func (i intAlias) testInt() {
	fmt.Println(i)
}

/**
* struct 作为函数参数,as a pointer
 */
func testYang(y *yang) {
	fmt.Println(y)
}

/**
* struct 作为函数参数，as a value type
 */
func testYangVal(y yang) {
	fmt.Println(y)
}

/**
* factory method
 */
func newYang() *yang {
	y := new(yang)
	return y
}

/**
* 工厂方法测试
 */
func factoryTest() {
	world := world.NewWorld()
	fmt.Println(world.Name, world.Size)
}

/**
* two yang method
 */
func (*twoYang) twoyangInfo() {
	fmt.Println("i am twoyang info")
}

func (*twoYang) echoInfo() {
	fmt.Println("two yang echo info method")
}

/**
* 匿名类型 测试
 */
func extendTest() {
	tt := new(twoYang)
	tt.name = "beijing"
	tt.age = 55
	tt.echoInfo()
	tt.twoyangInfo()
}
