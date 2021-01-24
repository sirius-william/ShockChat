package test

import (
	"ShockChatServer/utils"
)

func TestRSAEncrypt(str []byte) ([]byte, error) {
	res, err := utils.Encrypt(str, "public.pem")

	return res, err
}

func TestRSADecrypt(_byte []byte) string {
	res, err := utils.Decrypt(_byte, "private.pem")
	if err != nil {
		return ""
	} else {
		return string(res)
	}
}
