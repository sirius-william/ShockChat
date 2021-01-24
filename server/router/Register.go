// @Title  Register.go
// @Description	用户注册模块，消息id：
// @Author  SiriusWilliam  2021-01-22 9:50
// @Update  2021-01-22 9:50
package router

import (
	"ShockChatServer/protos"
	"ShockChatServer/utils"
	"fmt"
	"github.com/aceld/zinx/ziface"
	"github.com/aceld/zinx/znet"
	"github.com/fatih/color"
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
	registerInfo := protos.UserRegisterInfo{}
	status := protos.Status{}
	data := req.GetData()
	if len(data) == 0 {
		req.GetConnection().Stop()
		return
	}
	err := proto.Unmarshal(data, &registerInfo)
	if err != nil {
		status.Status = false
		status.Error = "server error."
		statusSend, _ := proto.Marshal(&status)
		_ = req.GetConnection().SendMsg(0x103, statusSend)
		color.Red("%s", err)
		return
	}
	userNameBytes, _ := utils.Decrypt(registerInfo.Name, utils.KeyFile.PrivateKeyFilePath)
	passwordBytes, _ := utils.Decrypt(registerInfo.Password, utils.KeyFile.PrivateKeyFilePath)
	telBytes, _ := utils.Decrypt(registerInfo.Tel, utils.KeyFile.PrivateKeyFilePath)
	emailBytes, _ := utils.Decrypt(registerInfo.Email, utils.KeyFile.PrivateKeyFilePath)
	userName := string(userNameBytes)
	password := string(passwordBytes)
	tel := string(telBytes)
	email := string(emailBytes)
	fmt.Println(userName, password, tel, email)
}
