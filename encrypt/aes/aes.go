package aes

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
)

type EncryptParam struct {
	Key  []byte
	Mode string
}

func (This *EncryptParam) Encrypt(origData []byte) ([]byte, error) {
	block, err := aes.NewCipher(This.Key)
	if err != nil {
		return nil, err
	}
	blockSize := block.BlockSize()
	if This.Mode == "ECB" {
		origData = ZeroPadding(origData, block.BlockSize())
	} else {
		origData = PKCS5Padding(origData, blockSize)
	}
	blockMode := cipher.NewCBCEncrypter(block, This.Key[:blockSize])
	crypted := make([]byte, len(origData))
	// 根据CryptBlocks方法的说明，如下方式初始化crypted也可以
	blockMode.CryptBlocks(crypted, origData)
	return crypted, nil
}

func (This *EncryptParam) Decrypt(crypted []byte) ([]byte, error) {
	block, err := aes.NewCipher(This.Key)
	if err != nil {
		return nil, err
	}
	blockSize := block.BlockSize()
	blockMode := cipher.NewCBCDecrypter(block, This.Key[:blockSize])
	origData := make([]byte, len(crypted))
	blockMode.CryptBlocks(origData, crypted)
	if This.Mode == "ECB" {
		origData = ZeroUnPadding(origData)
	} else {
		origData = PKCS5UnPadding(origData)
	}
	return origData, nil
}

func ZeroPadding(ciphertext []byte, blockSize int) []byte {
	padding := blockSize - len(ciphertext)%blockSize
	padtext := bytes.Repeat([]byte{0}, padding)
	return append(ciphertext, padtext...)
}

func ZeroUnPadding(origData []byte) []byte {
	length := len(origData)
	unpadding := int(origData[length-1])
	return origData[:(length - unpadding)]
}

func PKCS5Padding(ciphertext []byte, blockSize int) []byte {
	padding := blockSize - len(ciphertext)%blockSize
	padtext := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(ciphertext, padtext...)
}

func PKCS5UnPadding(origData []byte) []byte {
	length := len(origData)
	// 去掉最后一个字节 unpadding 次
	unpadding := int(origData[length-1])
	return origData[:(length - unpadding)]
}
