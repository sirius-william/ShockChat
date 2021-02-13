// @Title  GetFunc
// @Description
// @Author  SiriusWilliam  2021-02-12 22:31
// @Update  2021-02-12 22:31
package funcs

import (
	"ShockChatServer/utils"
	"errors"
	"strconv"
)

// 获取某些信息
func Get(key []string) ([]string, error) {
	if len(key) < 1 {
		return nil, errors.New("there is syntax error.")
	}
	//name := []string{"connect", "user", "server"}
	if key[0] == "connect" {
		if key[1] == "count" {
			return []string{strconv.Itoa(getConnectCount())}, nil
		} else {
			return nil, errors.New("there is syntax error.")
		}
	}
	return nil, errors.New("there is syntax error.")
}

func getConnectCount() int {
	return len(utils.ConnectionIdReflectorZinxConnID)
}
