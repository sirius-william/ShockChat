package test

import (
	"ShockChatServer/utils"
)

func TestRSAEncrypt(str []byte) ([]byte, error) {
	res, err := utils.Encrypt(str)

	return res, err
}

func TestRSADecrypt(_byte []byte) string {
	res, err := utils.Decrypt(_byte)
	if err != nil {
		return ""
	} else {
		return string(res)
	}
}
