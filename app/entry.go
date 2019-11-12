/**
* package main ，这样的包将会被编译安装到 gopath环境变量下的bin中，变成了可执行文件
* package 非main,这样的包将会被编译安装到 gopath环境变量下的pkg中，以.a结尾
 */
package main

import (
	"fmt"
	"go_study_pkg/greet"
	"go_study_pkg/greet/de"
)

func main() {
	fmt.Println(greet.Morning)
	fmt.Println(de.DeName)
}
