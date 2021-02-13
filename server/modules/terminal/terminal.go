// @Title  terminal.go
// @Description	控制台，用于监控服务器服务端
// @Author  SiriusWilliam  2021-01-22 17:54
// @Update  2021-01-22 17:54
package terminal

import (
	"github.com/aceld/zinx/ziface"
	"github.com/aceld/zinx/znet"
)

type Terminal struct {
	znet.BaseRouter
}

func (t *Terminal) Handle(request ziface.IRequest) {

}
