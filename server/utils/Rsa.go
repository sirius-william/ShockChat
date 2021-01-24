package utils

import (
	"bytes"
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"fmt"
	"os"
)

type KeyFileConfig struct {
	PublicKeyFilePath string `json:"publicKey"`
	PrivateKeyFilePath string `json:"privateKey"`
}

var KeyFile KeyFileConfig = KeyFileConfig{"public.pem", "private.pem"}

func Encrypt(str []byte, publicKeyName string)([]byte, error)  {
	file, err := os.Open(publicKeyName)
	if err != nil {
		fmt.Println("file:", err)
		return nil, err
	}
	defer file.Close()
	info, _ := file.Stat()
	size := info.Size()
	buff := make([]byte, size)
	_, _ = file.Read(buff)
	block, _ := pem.Decode(buff)
	publicKeyInterface, err := x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil{
		fmt.Println("parse key:", err)
		return nil, err
	}
	//类型断言
	publicKey := publicKeyInterface.(*rsa.PublicKey)
	////对明文进行加密
	var cipherText []byte
	keySize, srcSize := publicKey.Size(), len(str)
	//单次加密的长度需要减掉padding的长度，PKCS1为11
	offSet, once := 0, keySize-11
	buffer := bytes.Buffer{}
	for offSet < srcSize {
		endIndex := offSet + once
		if endIndex > srcSize {
			endIndex = srcSize
		}
		// 加密一部分
		bytesOnce, err := rsa.EncryptPKCS1v15(rand.Reader, publicKey, str[offSet:endIndex])
		if err != nil {
			return nil, err
		}
		buffer.Write(bytesOnce)
		offSet = endIndex
	}
	cipherText = buffer.Bytes()
	return cipherText, nil
}
func Decrypt(str []byte, privateKeyName string) ([]byte, error) {
	file,err:=os.Open(privateKeyName)
	if err!=nil{
		return nil, err
	}
	defer file.Close()
	//获取文件内容
	info, err := file.Stat()
	if err != nil {
		return nil, err
	}
	buf:=make([]byte,info.Size())
	_, _ = file.Read(buf)
	//pem解码
	block, _ := pem.Decode(buf)

	//X509解码
	privateKey, err := x509.ParsePKCS1PrivateKey(block.Bytes)
	if err!=nil{
		return nil, err
	}
	var plainText []byte
	keySize := privateKey.Size()
	srcSize := len(str)

	var offSet = 0
	var buffer = bytes.Buffer{}
	for offSet < srcSize {
		endIndex := offSet + keySize
		if endIndex > srcSize {
			endIndex = srcSize
		}
		bytesOnce, err := rsa.DecryptPKCS1v15(rand.Reader, privateKey, str[offSet:endIndex])
		if err != nil {
			return nil, err
		}
		buffer.Write(bytesOnce)
		offSet = endIndex
	}
	plainText = buffer.Bytes()
	return plainText, nil
}
