/**
* array and slice
* array,string 底层都是数组，都可应用slice
 */
package main

import (
	"fmt"
	"sort"
)

/**
* 数组学习
 */
func arrayStudy() {
	/*arr := [5]int{2, 3, 5}
	arr_str := [8]string{"zhangsan", "lisi"}
	fmt.Println(arr, arr_str)*/
	//appendArray()
	sortArray()
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
	/*sub_arr := arr[:]
	cap_num := cap(sub_arr)
	len_num := len(sub_arr)
	fmt.Println(cap_num, len_num)
	sum_num := sum(sub_arr)
	fmt.Println(sum_num)*/
	slice_arr := arr[1:3]
	fmt.Println(slice_arr)
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

/**
* for range
*
 */
func forRange() {
	items := []int{10, 20, 30, 40}
	for _, val := range items {
		val *= 2
	}
	fmt.Println(items)
}

/**
* copy 数组
 */
func copyArray() {
	src_arr := []int{1, 2, 3}
	//dst_arr := []int{4, 5, 6}
	dst_arr := make([]int, 10)
	copy(dst_arr, src_arr)
	fmt.Println(dst_arr)
}

/**
* append study
 */
func appendArray() {
	src_arr := []int{1, 3}
	result := append(src_arr, 4, 5, 6)
	fmt.Println(result)
}

/**
* 字符串在底层是一个byte array
*
 */
func stringArray() {
	//s := "zhangsan"
	//s_byte := []byte(s)
	//for _, val := range s_byte {
	//fmt.Printf("%c", val)
	//}
	to_byte := make([]byte, 10)
	copy(to_byte, "lisi")
	fmt.Println(to_byte)
}

/**
* byte study
 */
func stringByte() {
	//string and slice
	s := "zhangsan"
	ss := s[0:5]
	fmt.Println(ss)

	/*string change one char
	string to byte,
	change one char
	byte to string*/
	b_arr := []byte(s)
	b_arr[0] = 'Z'
	new_s := string(b_arr)
	fmt.Println(new_s)
}

/**
* 数组排序和查找
* search and sort
 */
func sortArray() {
	arr := []int{1, 6, 19, 4, 2, 9}
	str_arr := []string{"aha", "dalas", "cluo", "boluo"}
	sort.Ints(arr)
	fmt.Println(arr)
	sort.Strings(str_arr)
	fmt.Println(str_arr)
}

func test() {
	/*arr := map[string]int{"name": 2, "city": 3, "sex": 5}
	fmt.Println(arr)*/
	//第一种方式
	var citys map[string]string
	citys = make(map[string]string)
	citys["name"] = "zhangsan"
	citys["sex"] = "man"
	fmt.Println(citys)
	//第二种方式
	person := map[string]string{}
	person["sex"] = "women"
	person["eye"] = "double"

	arr := [3]int{}
	arr[0] = 1
	arr[1] = 2

	sliceA := arr[:]
	sliceA[0] = 100
	sliceA[1] = 8
	sliceA[2] = 5
	fmt.Println(arr, sliceA)

}

func main() {
	test()
}
