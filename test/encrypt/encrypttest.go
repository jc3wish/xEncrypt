package main

import (
	"encoding/base64"
	"fmt"
	"github.com/xEncrypt/encrypt/aes"
	"github.com/xEncrypt/encrypt/des"
	"github.com/xEncrypt/encrypt/des3"
)

type EncryptFun interface {
	Decrypt(crypted []byte) ([]byte, error)
	Encrypt(origData []byte) ([]byte, error)
}

var EncryptMap map[string]EncryptFun

func main() {
	desTest()
	des3Test()
	aesTest()
}

func desTest() {
	fmt.Println("des test start...")
	module := "des"
	EncryptMap = make(map[string]EncryptFun)
	EncryptMap[module] = &des.EncryptParam{Key: []byte("sfe023f_")}
	encrypts, err := EncryptMap[module].Encrypt([]byte("mytestname"))
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("encrypts:%s \r\n", base64.StdEncoding.EncodeToString(encrypts))

	data := base64.StdEncoding.EncodeToString(encrypts)
	d, err3 := base64.StdEncoding.DecodeString(data)
	if err3 != nil {
		fmt.Println(err3)
	}
	encrypts2, err2 := EncryptMap[module].Decrypt([]byte(d))
	if err2 != nil {
		fmt.Println(err2)
	}
	fmt.Printf("encrypts2:%s \r\n", string(encrypts2))
	fmt.Println("des test over...")
}

func des3Test() {
	fmt.Println("des3 test start...")
	module := "des3"
	EncryptMap = make(map[string]EncryptFun)
	EncryptMap[module] = &des3.EncryptParam{Key: []byte("sfe023f_sefiel#fi32lf3e!")}
	encrypts, err := EncryptMap[module].Encrypt([]byte("mytestname"))
	if err != nil {
		fmt.Println(err)
	}
	data := base64.StdEncoding.EncodeToString(encrypts)
	fmt.Printf("encrypts:%s \r\n", base64.StdEncoding.EncodeToString(encrypts))

	d, err3 := base64.StdEncoding.DecodeString(data)
	if err3 != nil {
		fmt.Println(err3)
	}
	encrypts2, err2 := EncryptMap[module].Decrypt([]byte(d))
	if err2 != nil {
		fmt.Println(err2)
	}
	fmt.Printf("encrypts2:%s \r\n", string(encrypts2))
	fmt.Println("des3 test over...")
}

func aesTest() {
	fmt.Println("aes test start...")
	module := "aes"
	EncryptMap = make(map[string]EncryptFun)
	EncryptMap[module] = &aes.EncryptParam{Key: []byte("sfe023f_9fd&fwfl")}
	encrypts, err := EncryptMap[module].Encrypt([]byte("mytestname"))
	if err != nil {
		fmt.Println(err)
	}
	data := base64.StdEncoding.EncodeToString(encrypts)
	fmt.Printf("encrypts:%s \r\n", base64.StdEncoding.EncodeToString(encrypts))

	d, err3 := base64.StdEncoding.DecodeString(data)
	if err3 != nil {
		fmt.Println(err3)
	}
	encrypts2, err2 := EncryptMap[module].Decrypt([]byte(d))
	if err2 != nil {
		fmt.Println(err2)
	}
	fmt.Printf("encrypts2:%s \r\n", string(encrypts2))
	fmt.Println("aes test over...")
}
