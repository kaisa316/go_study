/**
* 控制语句学习
 */
package main

import (
	"fmt"
)

/**
* if else  学习
 */
func ifStudy() {
	a := 5
	if a < 0 {
		fmt.Println(a, "less than zero:")
	} else if a < 10 {
		code, err := fmt.Println(a, "between 0-10")
		if err != nil {
			fmt.Println(code, err)
		}
	} else {
		fmt.Println(a, "grater than 10")
	}
}

/**
* if语句中定义变量
 */
func ifSetVal() {
	//b的作用域仅在if语句中
	if b := 3; b < 5 {
		fmt.Println("b小于5")
	}
	//fmt.Println(b) //b在这里不可见
}

/**
* for 学习
 */
func forStudy() {
	//计数器
	/*for i := 0; i < 5; i++ {
		fmt.Println(i)
	}*/
	//条件
	/*i := 0
	for i <= 6 {
		i++
		if i == 3 {
			continue
		}
		fmt.Println(i)
	}*/
	/*j := 1
	for {
		fmt.Println(j)
		j--
		if j <= -10 {
			break
		}
	}
	*/
outline:
	for i := 1; i <= 5; i++ {
		for j := 1; j <= 3; j++ {
			if j == 2 {
				//break
				//continue
				break outline
			}
			fmt.Println(i * j)
		}
	}
	fmt.Println("end")
	//for range
	/*s := "中华民族"
	for ix, val := range s {
		//fmt.Println(ix, val)
		fmt.Printf("%d,%d,%c\n", ix, val, val)
	}*/

}

/**
* goto 用于跳跃多层循环
 */
func forGoto() {
	//如果是跳转到这里，则成了无限循环
	//TagTwo:
	for i := 5; i < 10; i++ {
		for j := 0; j < 3; j++ {
			if i == 6 {
				goto TagOne
			}
			fmt.Println(i, ",", j)
		}
	}
TagOne:
	//循环跳到这里
}

func forBreakContinue() {
	/*
		for i := 0; i < 5; i++ {
			if i == 3 {
				break
			}
			fmt.Println("I is ", i)
		}

		fmt.Println("----------------")
		for j := 0; j < 10; j++ {
			if j == 5 {
				continue //跳过当次循环
			}
			fmt.Println("j is ", j)
		}
	*/

	for i := 0; i < 5; i++ {
		for j := 1; j < 5; j++ {
			if i == 2 {
				break //仅跳出j层循环
			}
			fmt.Println(i, ",", j)
		}
	}
}

/**
* switch 学习
 */
func switchStudy() {
	exp := 5
	switch exp {
	case 4:
		fmt.Println("equal 4")
	case 5:
		fmt.Println("equal 5")
		fallthrough
	case 6:
		fmt.Println("equal 6")
	default:
		fmt.Println("is default")
	}

	j := 'c'
	switch j {
	case 'a', 'c':
		fmt.Println("a or c")
	case 'b':
		fmt.Println("is b")
	default:
		fmt.Println(" is default ")
	}

}

func test() {
	for i := 0; i < 5; i++ {
		fmt.Println(i)
	}

	j := 10
	for j < 5 {
		fmt.Println(j)
		j--
	}

}

func main() {
	forGoto()
}
