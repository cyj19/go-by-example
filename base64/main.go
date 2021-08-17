package main

import (
	"encoding/base64"
	"fmt"
)

func main() {
	data := "abc123!?$*&()'-=@~"

	//编码
	sEnc := base64.StdEncoding.EncodeToString([]byte(data))
	fmt.Println(sEnc)

	//解码
	sDec, _ := base64.RawStdEncoding.DecodeString(sEnc)
	fmt.Println(string(sDec))
	fmt.Println()

	uEnc := base64.URLEncoding.EncodeToString([]byte(data))
	fmt.Println(uEnc)
	uDec, _ := base64.RawURLEncoding.DecodeString(uEnc)
	fmt.Println(string(uDec))

}
