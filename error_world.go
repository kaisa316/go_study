package main

import "errors"
import "fmt"

func main() {
	//check()
	exceptionName("")
}

//不会终止程序的执行，仅作为return value形式返回
//多用于逻辑不正确时的错误展示
func check() {
	name := ""
	err := checkName(name)
	if err != nil {
		fmt.Println("发生了错误:", err)
	}

}

func checkName(name string) error {
	if name == "" {
		return errors.New("name is empty")
	}
	return nil
}

func exceptionName(name string) int {
	defer catch()
	if name == "" {
		panic("发生异常，终止执行")
	}
	return 1
}

func catch() {
	fmt.Println("catch error")
}
