// @Title  LegalCheck.go
// @Description	验证连接合法性模块，消息id：0x100，0x101,0x102
// @Author  SiriusWilliam  2021-01-21 15:43
// @Update  2021-01-21 15:43
package router

import (
	"ShockChatServer/logger"
	"ShockChatServer/protos"
	"ShockChatServer/utils"
	"fmt"
	"github.com/aceld/zinx/ziface"
	"github.com/aceld/zinx/znet"
	"github.com/golang/protobuf/proto"
)

/*
 * id: 0x100
 * 描述：客户端请求验证连接合法性
 */
type LegalCheckSaltRouter struct {
	znet.BaseRouter
}

func (br *LegalCheckSaltRouter) Handle(req ziface.IRequest) {
	var data []byte = req.GetData()
	// 这个消息必须为空
	if len(data) != 0 {
		return
	}
	// 拿到盐值
	var salt string = utils.GetSalt()
	// 计算谜语并写入属性
	utils.Riddle(salt, req)
	// 写入响应体
	var respond protos.LegalCheckSalt = protos.LegalCheckSalt{Salt: salt}
	var err error
	var send []byte
	// 序列化
	send, err = proto.Marshal(&respond)
	if err != nil {
		// 发生错误，响应server error
		var status protos.Status = protos.Status{}
		status.Status = false
		status.Message = "server error"
		logger.Log.Error(err.Error())
		send, _ = proto.Marshal(&status)
		err = req.GetConnection().SendMsg(0x000, send)
		if err != nil {
			logger.Log.Error(err.Error())
		}
		return
	}
	// 发送盐值响应体给客户端
	err = req.GetConnection().SendMsg(0x101, send)
	if err != nil {
		logger.Log.Error(err.Error())
	}
}

/*
 *	id:0x102
 */
type SaltCheck struct {
	znet.BaseRouter
}

func (br *SaltCheck) Handle(req ziface.IRequest) {
	var data []byte = req.GetData()
	if len(data) == 0 {
		return
	}
	var respondFromUser protos.LegalCheckResult
	var err error
	var send []byte
	var checkRes protos.LegalCheckStatus
	err = proto.Unmarshal(data, &respondFromUser)
	if err != nil {
		checkRes.Status = -1
		checkRes.Error = "server error"
		logger.Log.Error(err.Error())
		send, err = proto.Marshal(&checkRes)
		if err != nil {
			logger.Log.Error(err.Error())
			return
		}
		err = req.GetConnection().SendMsg(0x103, send)
		if err != nil {
			logger.Log.Error(err.Error())
			return
		}
	}
	// 拿到之前存储的计算好的谜题结果
	riddle, err := req.GetConnection().GetProperty("riddle")
	if err != nil {
		logger.Log.Error(err.Error())
		return
	}
	// 对比
	if riddle.(string) == respondFromUser.Result {
		checkRes.Status = 1
		checkRes.Error = ""
		req.GetConnection().SetProperty("checked", true)
	} else {
		checkRes.Status = 0
		checkRes.Error = "check incorrect"
	}
	send, err = proto.Marshal(&checkRes)
	if err != nil {
		fmt.Println("序列化错误")
		logger.Log.Error(err.Error())
		return
	}
	err = req.GetConnection().SendMsg(0x103, send)
	req.GetConnection().RemoveProperty("salt")
}

func IsLegal(req ziface.IRequest) bool {
	checked, err := req.GetConnection().GetProperty("checked")
	if checked != nil && err == nil {
		if checked.(bool) == true {
			fmt.Println("合法连接")
			return true
		}
	}
	return false
}
