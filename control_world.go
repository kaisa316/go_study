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
* switch 学习
 */
func switchStudy() {
	i := 'c'
	switch i {
	case 'a', 'c':
		fmt.Println("china bisheng")
	case 'b':
		fmt.Println("american bisheng")
	default:
		fmt.Println("default")

	}

}
