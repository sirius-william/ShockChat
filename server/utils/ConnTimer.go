// @Title  ConnTimer.go
// @Description
// @Author  SiriusWilliam  2021-01-26 20:01
// @Update  2021-01-26 20:01
package utils

import (
	"ShockChatServer/protos"
	"ShockChatServer/utils/redis"
	"fmt"
	"github.com/aceld/zinx/ziface"
	"time"
)

var TimerPool map[uint32]chan bool = make(map[uint32]chan bool)

/*
 * 根据Token配置，在Token过期前一分钟定时将Token发送给客户端，这个方法在用户登录后执行
 */
func StartTimerOfSendingToken(user *protos.UserLogin, conn ziface.IConnection) {
	TimerPool[conn.GetConnID()] = make(chan bool)
	go func() {
		timer := time.NewTicker(time.Duration(TokenConf.Minute-1) * time.Minute)
		for {
			select {
			case <-TimerPool[conn.GetConnID()]:
				{
					timer.Stop()
					return
				}
			case <-timer.C:
				{
					token := CreateToken(user)
					// 将新生成的token写入redis数据库
					_, _ = redis.Exec("hset", "tokens", user.Id, token)
					err := conn.SendMsg(0x301, []byte(token))
					if err != nil {
						fmt.Println(err)
					}
				}
			}
		}
	}()
}
