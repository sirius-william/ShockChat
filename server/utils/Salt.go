// @Title  Riddle
// @Description	用于验证合法性的猜谜游戏
// @Author  SiriusWilliam  2021-01-21 21:53
// @Update  2021-01-21 21:53
package utils

import (
	"github.com/aceld/zinx/ziface"
	"math/rand"
	"strconv"
	"time"
)

/*
 * 服务端计算的猜谜，算法目前由于测试，规则为计算md5值，可根据需要进行修改，这个规则要与客户端相同
 */
func Riddle(salt string, req ziface.IRequest) {
	// 计算出md5值
	m := StringToMD5(salt)
	// 将md5写入连接的自定义属性中
	req.GetConnection().SetProperty("legal_check", m)
}

func GetSalt() string {
	rand.Seed(time.Now().Unix())
	var salt int = rand.Intn(99999999) + 10000000
	return strconv.Itoa(salt)
}
