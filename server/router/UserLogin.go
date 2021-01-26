// @Title  UserLogin.go
// @Description	用户登录路由
// @Author  SiriusWilliam  2021-01-26 21:11
// @Update  2021-01-26 21:11
package router

import (
	"github.com/aceld/zinx/ziface"
	"github.com/aceld/zinx/znet"
)

type LoginRouter struct {
	znet.BaseRouter
}

func (br *LoginRouter) Handle(req ziface.IRequest) {

}
