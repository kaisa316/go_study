package main

import (
	"bytes"
	"encoding/json"
	"fmt"

	jsoniter "github.com/json-iterator/go"
	"github.com/techleeone/gophp/serialize"
)

func main() {
	// guanfangInterface()
	// jsonIteratorInterface()
	// seriInterface()
	gobTest()
}

func guanfang() {
	s := map[string]string{
		"name": "zhangsan",
		"age":  "23",
	}

	result, err := json.Marshal(s)
	ss := string(result)
	fmt.Println(string(result))

	var parse interface{}
	err = json.Unmarshal([]byte(ss), &parse)
	//a := parse.(map[string]string)
	fmt.Sprintf("%v,%v,%T", parse, err)

}

func guanfangInterface() {
	s := map[string]interface{}{
		"name": "zhangsan",
		"age":  23,
	}

	result, _ := json.Marshal(s)
	ss := string(result)
	fmt.Println(string(result))

	var parse interface{}
	json.Unmarshal([]byte(ss), &parse)
	a := parse.(map[string]interface{})

	fmt.Printf("官网输出类型：%T", a["age"])

}

func jsonIteratorInterface() {
	s := map[string]interface{}{
		"name":   "zhangsan",
		"age":    23,
		"weight": 99.45,
	}

	// result, err := json.Marshal(s)
	result, _ := jsoniter.Marshal(s)
	ss := string(result)
	// fmt.Println(string(result))
	fmt.Println()

	var parse interface{}
	// err = json.Unmarshal([]byte(ss), &parse)
	// err := jsoniter.Unmarshal([]byte(ss), &parse)
	// a := parse.(map[string]interface{})
	decoder := json.NewDecoder(bytes.NewReader([]byte(ss)))
	decoder.UseNumber()
	decoder.Decode(&parse)
	a := parse.(map[string]interface{})
	fmt.Printf("第三方输出类型 :age:%T,weight:%T", a["age"], a["weight"])

}

//序列化
func seriInterface() {
	// s := map[string]interface{}{
	// 	"name":   "zhangsan",
	// 	"age":    23,
	// 	"weight": 99.45,
	// }

	type sa struct {
		name   string
		age    int
		weight float64
	}

	saval := sa{
		name:   "zhangsan",
		age:    23,
		weight: 99.45,
	}

	var json = jsoniter.ConfigCompatibleWithStandardLibrary
	ss, _ := json.MarshalToString(saval)
	fmt.Println(ss)
	var data sa
	json.UnmarshalFromString(ss, &data)
	fmt.Printf("%v,%T", data, data.age)

}

func gobTest() {
	s := map[string]interface{}{
		"name":   "zhangsan",
		"age":    23,
		"weight": 99.45,
	}

	//encode
	bytes, _ := serialize.Marshal(s)
	ss := string(bytes)
	fmt.Printf("%v,%T", ss, ss)

	//decode
	decode, _ := serialize.UnMarshal([]byte(ss))
	decodeMap := decode.(map[string]interface{})
	fmt.Printf("%v,%T", decodeMap, decodeMap["age"])

}
