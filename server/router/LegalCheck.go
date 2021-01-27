// @Title  LegalCheck.go
// @Description	验证连接合法性模块，消息id：0x100，0x101,0x102
// @Author  SiriusWilliam  2021-01-21 15:43
// @Update  2021-01-21 15:43
package router

import (
	"ShockChatServer/protos"
	"ShockChatServer/utils"
	"github.com/aceld/zinx/ziface"
	"github.com/aceld/zinx/znet"
	"github.com/golang/protobuf/proto"
	"math/rand"
	"strconv"
	"time"
)

/*
 * id: 0x100
 */
type LegalCheckSaltRouter struct {
	znet.BaseRouter
}

func (br *LegalCheckSaltRouter) Handle(req ziface.IRequest) {
	// 验证消息格式，此消息只包含报头，没有包体
	if len(req.GetData()) != 0 {
		err := protos.Status{Status: false, Error: "invalid request"}
		errSend, _ := proto.Marshal(&err)
		_ = req.GetConnection().SendMsg(0x103, errSend)
		return
	}
	// 重置随机数种子，后获取一个随机1000~2000整数
	rand.Seed(time.Now().Unix())
	salt := rand.Intn(9999) + 1000
	var saltSend = protos.LegalCheckSalt{Salt: strconv.Itoa(salt)}
	saltByte, _ := proto.Marshal(&saltSend)
	//saltRSA, _ := RSA.Encrypt(saltByte, "public.pem")
	utils.Riddle(salt, req)
	_ = req.GetConnection().SendMsg(0x101, saltByte)
}

/*
 *	id:0x102
 */
type SaltCheck struct {
	znet.BaseRouter
}

func (br *SaltCheck) Handle(req ziface.IRequest) {
	res := protos.Status{}
	var resSend []byte
	// 拿到数据包
	data := req.GetData()
	// 反序列化拿到客户端猜谜结果
	salt := protos.LegalCheckResult{}
	err := proto.Unmarshal(data, &salt)
	// 是否解析成功，未成功则返回
	if err != nil {
		res.Status = false
		res.Error = "invalid request"
		resSend, _ = proto.Marshal(&res)
		_ = req.GetConnection().SendMsg(0x103, resSend)
		return
	}
	// 拿到之前服务器这边的猜谜结果
	serverCalculated, err := req.GetConnection().GetProperty("legal_check")
	// 如果：1）发现没有猜谜，即发现并没有之前的获取盐值操作，返回错误的结果给客户端
	// 2）发现对比正确，返回正确给客户端，并将checked=true写入连接属性，用于后续通信，也防止了反复验证
	// 3) 猜谜错误，视为非法连接，将连接关闭并清除
	if err != nil {
		res.Status = false
		res.Error = "server error"
	} else if serverCalculated.(string) == salt.GetResult() {
		res.Status = true
		res.Error = ""
		req.GetConnection().SetProperty("checked", true)
	} else {
		res.Status = false
		res.Error = serverCalculated.(string)
		req.GetConnection().Stop()
	}
	resSend, _ = proto.Marshal(&res)
	_ = req.GetConnection().SendMsg(0x103, resSend)
}

/*
 * 合法性检查，如果非法，关闭连接，返回false，该函数执行后，若为false，一般禁止对连接进行再次操作，因为连接关闭了。
 */
func LegalCheck(req ziface.IRequest) bool {
	legal, err := req.GetConnection().GetProperty("checked")
	if err != nil || legal.(bool) != true {
		// 关闭非法连接
		_ = req.GetConnection().SendMsg(0x103, []byte("illegal connection"))
		req.GetConnection().Stop()
		return false
	}
	return true
}
