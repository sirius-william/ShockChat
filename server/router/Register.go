// @Title  Register.go
// @Description	用户注册模块，消息id：
// @Author  SiriusWilliam  2021-01-22 9:50
// @Update  2021-01-22 9:50
package router

import (
	"ShockChatServer/logger"
	"ShockChatServer/protos"
	"ShockChatServer/utils"
	"ShockChatServer/utils/mysql"
	"github.com/aceld/zinx/ziface"
	"github.com/aceld/zinx/znet"
	"github.com/golang/protobuf/proto"
)

/*
 * id:0x200
 * 用户申请注册账户
 */
type RegisterRouter struct {
	znet.BaseRouter
}

func (r *RegisterRouter) Handle(req ziface.IRequest) {
	if !IsLegal(req) {
		// 非法连接，断开
		req.GetConnection().Stop()
		return
	}
	registerInfo := protos.UserRegisterInfo{}
	data := req.GetData()
	if len(data) == 0 {
		req.GetConnection().Stop()
		return
	}
	// 解析
	err := proto.Unmarshal(data, &registerInfo)
	if err != nil {
		logger.Log.Error(err.Error())
		return
	}
	userNameBytes, _ := utils.Decrypt(registerInfo.Name)
	passwordBytes, _ := utils.Decrypt(registerInfo.Password)
	telBytes, _ := utils.Decrypt(registerInfo.Tel)
	emailBytes, _ := utils.Decrypt(registerInfo.Email)
	username := string(userNameBytes)
	password := string(passwordBytes)
	tel := string(telBytes)
	email := string(emailBytes)
	// 用户id
	var id int
	userid := protos.UserId{Id: int64(id), Error: ""}
	// 操作数据库
	id, err = mysql.Register(username, password, tel, email)
	if err != nil {
		userid.Id = -1
		userid.Error = "server error"
		logger.Log.Error(err.Error())
	}
	send, err := proto.Marshal(&userid)
	if err != nil {
		logger.Log.Error(err.Error())
		return
	}
	_ = req.GetConnection().SendMsg(0x201, send)
}
