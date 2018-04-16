package main

import (
	"fmt"
	"github.com/xEncrypt/encrypt"
)

func main() {
	DoEncrypt()
	DoDecrypt()
}

func DoEncrypt() {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)
		}
	}()
	data, err := encrypt.DoEncrypt("email2", "mytestname")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(data)
}

func DoDecrypt() {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)
		}
	}()
	d := "N1/u1DamOOOHzxOExyJHQA=="
	//d := "DG3CgDUmIT/IENhFtsYreA=="
	data, err := encrypt.DoDecrypt("email2", d)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(data)
}
