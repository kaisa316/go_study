/**
* map 学习
 */
package main

import (
	"fmt"
	"go_study_pkg"
)

func mapStudy() {
	/*var test = make(map[string]string) //初始化
	test["name"] = "zhangsan"
	test["age"] = "23"
	test2 := map[string]int{"name": 21, "age": 33}

	var test3 map[string]string
	test3 = map[string]string{"name": "wang"} //没有这行，下面的赋值是不行的
	test3["age"] = "33"
	test3["ageKdd"] = "33"
	fmt.Println(test, test2, test3)*/
	pkg_test := go_study_pkg.Hello_pkg()
	fmt.Println(pkg_test)

}
