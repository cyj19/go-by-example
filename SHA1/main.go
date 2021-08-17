package main

import (
	"crypto/sha1"
	"fmt"
)

func main() {
	s := "sha1 this string"

	//返回一个新的使用SHA1校验的hash.Hash接口
	h := sha1.New()
	//将s进行散列
	h.Write([]byte(s))
	//返回添加b到当前的hash值后的新切片，不会改变底层的hash状态
	bs := h.Sum(nil)

	fmt.Println(s)
	fmt.Printf("%x\n", bs)

}
