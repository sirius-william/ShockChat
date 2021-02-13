// @Title  UserLoginDB.go
// @Description	用户登录路由
// @Author  SiriusWilliam  2021-01-26 21:11
// @Update  2021-01-26 21:11
package router

import (
	"ShockChatServer/logger"
	"ShockChatServer/protos"
	"ShockChatServer/utils"
	"ShockChatServer/utils/mysql"
	"fmt"
	"github.com/aceld/zinx/ziface"
	"github.com/aceld/zinx/znet"
	"github.com/golang/protobuf/proto"
)

type LoginRouter struct {
	znet.BaseRouter
}

func (br *LoginRouter) Handle(req ziface.IRequest) {
	isReqChecked := IsLegal(req)
	if isReqChecked == false {
		req.GetConnection().Stop()
		return
	}
	var data []byte = req.GetData()
	var userLoginInfo protos.UserLogin
	var err error
	var status protos.LoginResult
	var send []byte
	if len(data) == 0 {
		return
	}
	err = proto.Unmarshal(data, &userLoginInfo)
	if err != nil {
		status.IsSuccess = false
		status.Status = -1
		status.Error = "server error"
		send, err = proto.Marshal(&status)
		if err != nil {
			logger.Log.Error(err.Error())
			return
		}
		err = req.GetConnection().SendMsg(0x301, send)
		if err != nil {
			logger.Log.Error(err.Error())
			return
		}
	}
	var userid int32 = userLoginInfo.Id
	var passwordBytes []byte
	var password string
	fmt.Println(len(userLoginInfo.Password))
	passwordBytes, err = utils.Decrypt(userLoginInfo.Password)
	if err != nil {
		logger.Log.Error(err.Error())
		return
	}
	password = string(passwordBytes)
	var isChecked bool
	isChecked, err = mysql.UserLogin(userid, password)
	status.Status = 0
	status.IsSuccess = isChecked
	status.Error = err.Error()
	var token string
	if isChecked {
		token, err = utils.CreateToken(&userLoginInfo)
		if err != nil {
			logger.Log.Error(err.Error())
			return
		}
		status.Error = token
	}
	send, _ = proto.Marshal(&status)
	err = req.GetConnection().SendMsg(0x301, send)
	if err != nil {
		logger.Log.Error(err.Error())
		return
	}
	// 将连接存放至连接池内
	utils.ConnectionIdReflectorZinxConnID.Add(userid, req.GetConnection())
}
