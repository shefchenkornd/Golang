package main

import (
	"encoding/base64"
	"fmt"
)

func main() {
	data := "Hello world"

	sEnc := base64.StdEncoding.EncodeToString([]byte(data))
	fmt.Println(sEnc) // SGVsbG8gd29ybGQ=

	sDec, _ := base64.StdEncoding.DecodeString(sEnc)
	fmt.Println(string(sDec)) // Hello world

	uEnc := base64.URLEncoding.EncodeToString([]byte(data))
	fmt.Println(uEnc) // SGVsbG8gd29ybGQ=

	uDec, _ := base64.URLEncoding.DecodeString(uEnc)
	fmt.Println(string(uDec)) // Hello world
}
