// @Title  SendMessageMgDB.go
// @Description
// @Author  SiriusWilliam  2021-02-03 10:42
// @Update  2021-02-03 10:42
package mongodb

import (
	"ShockChatServer/protos"
	"time"
)

/*
 * 发送消息
 */
func SendMessage(fromId int32, toId int32, msg string, time time.Duration) protos.Messages {
	var MessageRes protos.Messages
	//append(MessageRes.Msg, &protos.Message{Msg: []byte(msg), Userid: 100})
	return MessageRes
}
