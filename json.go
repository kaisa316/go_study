package main

import (
	"bytes"
	"encoding/json"
	"fmt"

	jsoniter "github.com/json-iterator/go"
)

func main() {
	guanfangInterface()
	jsonIteratorInterface()
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
