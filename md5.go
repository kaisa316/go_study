package main

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
)

func main() {
	//m := md5.New()
	//data := []byte("Mdroid.cn")
	//s := md5.Sum(data)
	//ss := hex.EncodeToString(s)
	//fmt.Println(s)
	//fmt.Println(ss)
	md5Test()
}

func md5Test() {
	data := []byte("zhangsan")
	md5Ctx := md5.New()
	md5Ctx.Write(data)
	cipherStr := md5Ctx.Sum(nil)
	fmt.Println(cipherStr)
	fmt.Printf("%x\n", md5.Sum(data))
	fmt.Printf("%x\n", cipherStr)
	fmt.Println(hex.EncodeToString(cipherStr))
	fmt.Println("php result :", "01d7f40760960e7bd9443513f22ab9af")
}
