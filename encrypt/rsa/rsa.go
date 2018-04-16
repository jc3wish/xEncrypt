package rsa

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"errors"
)

type EncryptParam struct {
	PrivateKey []byte
	PublicKey  []byte
}

// 加密
func (This *EncryptParam) Encrypt(origData []byte) ([]byte, error) {
	block, _ := pem.Decode(This.PublicKey)
	if block == nil {
		return nil, errors.New("public key error")
	}
	pubInterface, err := x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		return nil, err
	}
	pub := pubInterface.(*rsa.PublicKey)
	return rsa.EncryptPKCS1v15(rand.Reader, pub, origData)
}

// 解密
func (This *EncryptParam) Decrypt(ciphertext []byte) ([]byte, error) {
	block, _ := pem.Decode(This.PrivateKey)
	if block == nil {
		return nil, errors.New("private key error!")
	}
	priv, err := x509.ParsePKCS1PrivateKey(block.Bytes)
	if err != nil {
		return nil, err
	}
	return rsa.DecryptPKCS1v15(rand.Reader, priv, ciphertext)
}
