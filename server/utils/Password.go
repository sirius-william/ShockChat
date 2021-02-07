// @Title  Password.go
// @Description 密码加密算法
// @Author  SiriusWilliam  2021-02-04 20:57
// @Update  2021-02-04 20:57
package utils

func Password(password string) string {
	return StringToMD5(password)
}
