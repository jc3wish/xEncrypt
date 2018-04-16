package encrypt

import (
	"encoding/base64"
	"errors"
	//"fmt"
	"github.com/xEncrypt/conf"
	"github.com/xEncrypt/encrypt/aes"
	"github.com/xEncrypt/encrypt/des"
	"github.com/xEncrypt/encrypt/des3"
	"github.com/xEncrypt/encrypt/rsa"
)

type EncryptFun interface {
	Decrypt(crypted []byte) ([]byte, error)
	Encrypt(origData []byte) ([]byte, error)
}

var EncryptMap map[string]EncryptFun

func init() {
	EncryptMap = make(map[string]EncryptFun)
	configInit()
}

func MyConfPanic(key string) {
	panic(key + " Mode must be CBC or ECB ")
}

func configInit() {
	for key, val := range conf.GetMyConf() {
		switch val.Type {
		case "aes":
			if val.Mode != "CBC" && val.Mode != "ECB" {
				MyConfPanic(key)
			}
			if len([]byte(val.Key)) != 16 {
				panic(key + " aes key len must be 16 ")
			}
			EncryptMap[key] = &aes.EncryptParam{Key: []byte(val.Key), Mode: val.Mode}
		case "des":
			if val.Mode != "CBC" && val.Mode != "ECB" {
				MyConfPanic(key)
			}
			if len([]byte(val.Key)) != 8 {
				panic(key + "  des key len must be 8 ")
			}
			EncryptMap[key] = &des.EncryptParam{Key: []byte(val.Key), Mode: val.Mode}
		case "des3":
			fallthrough
		case "3des":
			if val.Mode != "CBC" && val.Mode != "ECB" {
				MyConfPanic(key)
			}
			if len([]byte(val.Key)) != 24 {
				panic(key + " des3 key len must be 24 ")
			}
			EncryptMap[key] = &des3.EncryptParam{Key: []byte(val.Key), Mode: val.Mode}
		case "rsa":
			EncryptMap[key] = &rsa.EncryptParam{PrivateKey: []byte(val.PrivateKey), PublicKey: []byte(val.PublicKey)}
		default:
			continue
		}
	}
}

func DoEncrypt(key string, val string) (string, error) {
	Fun, ok := EncryptMap[key]
	if !ok {
		return "", errors.New("no key:" + key + " config")
	}
	data, err := Fun.Encrypt([]byte(val))
	if err != nil {
		return "", err
	}
	return base64.StdEncoding.EncodeToString(data), nil
}

func DoDecrypt(key string, val string) (string, error) {
	Fun, ok := EncryptMap[key]
	if !ok {
		return "", errors.New("no key:" + key + " config")
	}
	d, err3 := base64.StdEncoding.DecodeString(val)
	if err3 != nil {
		return "", err3
	}
	data, err := Fun.Decrypt([]byte(d))
	return string(data), err
}
