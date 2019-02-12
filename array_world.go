package main

import (
	"fmt"
)

/**
* 数组学习
 */
func arrayStudy() {
	arr := [5]int{2, 3, 5}
	arr_str := [8]string{"zhangsan", "lisi"}
	fmt.Println(arr, arr_str)
}

/**
*
* slice 切片学习
* slice ，即数组的一部分
*
 */
func sliceStudy() {
	//sub_arr := []string
	arr := [5]int{2, 1, 3, 5, 7}
	//var sub_arr []int
	//sub_arr := arr[1:4]
	//sub_arr := arr[:3]
	//sub_arr := arr[2:]
	sub_arr := arr[:]
	cap_num := cap(sub_arr)
	len_num := len(sub_arr)
	fmt.Println(cap_num, len_num)
	sum_num := sum(sub_arr)
	fmt.Println(sum_num)
}

func sum(arr []int) int {
	sum_num := 0
	for _, val := range arr {
		sum_num += val
	}
	//for i := 0; i < len(arr); i++ {
	//sum_num += arr[i]
	//}
	return sum_num
}
