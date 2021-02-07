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
	"github.com/aceld/zinx/ziface"
	"github.com/aceld/zinx/znet"
	"github.com/golang/protobuf/proto"
)

type LoginRouter struct {
	znet.BaseRouter
}

func (br *LoginRouter) Handle(req ziface.IRequest) {
	if IsLegal(req) == false {
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
	passwordBytes, err = utils.Decrypt(userLoginInfo.Password, utils.KeyFile.PrivateKeyFilePath)
	if err != nil {
		logger.Log.Error(err.Error())
		return
	}
	password = string(passwordBytes)
	var isChecked bool
	isChecked, err = mysql.UserLogin(userid, password)
	if err != nil {
		status.Status = 0
		status.IsSuccess = isChecked
		if err.Error() == "userid or password is incorrect" {
			status.Error = "userid or password is incorrect"
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
		} else {
			logger.Log.Error(err.Error())
			return
		}
	}
}
