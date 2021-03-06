package des3

import (
	"bytes"
	"crypto/cipher"
	"crypto/des"
)

type EncryptParam struct {
	Key  []byte
	Mode string
}

// 3DES加密
func (This *EncryptParam) Encrypt(origData []byte) ([]byte, error) {
	key := This.Key
	block, err := des.NewTripleDESCipher(key)
	if err != nil {
		return nil, err
	}
	if This.Mode == "ECB" {
		origData = ZeroPadding(origData, block.BlockSize())
	} else {
		origData = PKCS5Padding(origData, block.BlockSize())
	}
	blockMode := cipher.NewCBCEncrypter(block, key[:8])
	crypted := make([]byte, len(origData))
	blockMode.CryptBlocks(crypted, origData)
	return crypted, nil
}

// 3DES解密
func (This *EncryptParam) Decrypt(crypted []byte) ([]byte, error) {
	key := This.Key
	block, err := des.NewTripleDESCipher(key)
	if err != nil {
		return nil, err
	}
	blockMode := cipher.NewCBCDecrypter(block, key[:8])
	origData := make([]byte, len(crypted))
	// origData := crypted
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
	return bytes.TrimRightFunc(origData, func(r rune) bool {
		return r == rune(0)
	})
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
	l := length - unpadding
	return origData[:l]
}
