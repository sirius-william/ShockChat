// @Title  md5.go
// @Description	生成md5
// @Author  SiriusWilliam  2021-01-25 12:28
// @Update  2021-01-25 12:28
package utils

import (
	md52 "crypto/md5"
	"encoding/hex"
)

func StringToMD5(str string) string {
	md5 := md52.New()
	md5.Write([]byte(str))
	res := md5.Sum(nil)
	return hex.EncodeToString(res)
}
