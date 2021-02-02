// @Title  UserLogin.go
// @Description	用户登录路由
// @Author  SiriusWilliam  2021-01-26 21:11
// @Update  2021-01-26 21:11
package router

import (
	"ShockChatServer/protos"
	"ShockChatServer/utils"
	"ShockChatServer/utils/mysql"
	"github.com/aceld/zinx/ziface"
	"github.com/aceld/zinx/znet"
	"github.com/fatih/color"
	"github.com/golang/protobuf/proto"
)

type LoginRouter struct {
	znet.BaseRouter
}

func (br *LoginRouter) Handle(req ziface.IRequest) {
	if LegalCheck(req) == false {
		return
	}
	data := req.GetData()
	var userLoginInfo protos.UserLogin
	var status protos.Status
	var statusSend []byte
	err := proto.Unmarshal(data, &userLoginInfo)
	if err != nil {
		_ = req.GetConnection().SendMsg(0x103, []byte("server error"))
		return
	}
	// 拿到id
	id := userLoginInfo.Id
	// 检查用户是否已经登录
	_, ok := utils.ConnectionIdReflectorZinxConnID[id]
	if ok {
		status.Error = "you have login."
		status.Status = false
		statusSend, _ = proto.Marshal(&status)
		_ = req.GetConnection().SendMsg(0x103, statusSend)
		return
	}
	// 拿到密码，并解密
	password, err := utils.Decrypt(userLoginInfo.Password, utils.KeyFile.PrivateKeyFilePath)
	// 计算md5
	pwdMd5 := utils.StringToMD5(string(password))
	// 开启MySQL事务
	tx, err := mysql.Mysql.Begin()
	/*
		 * SQL说明：由MySQL评估，该查询语句策略为：
		+----+-------------+---------------+------------+-------+---------------+---------+---------+-------+------+----------+-------+
		| id | select_type | table         | partitions | type  | possible_keys | key     | key_len | ref   | rows | filtered | Extra |
		+----+-------------+---------------+------------+-------+---------------+---------+---------+-------+------+----------+-------+
		|  1 | SIMPLE      | tb_user_login | NULL       | const | PRIMARY       | PRIMARY | 4       | const |    1 |   100.00 | NULL  |
		+----+-------------+---------------+------------+-------+---------------+---------+---------+-------+------+----------+-------+
		由上，该查询语句走主键聚簇索引，性能卓越
	*/
	rows, err := tx.Query("select userid, password from tb_user_login where userid = ?;", id)
	var resId int
	var resPassword string
	for rows.Next() {
		err = rows.Scan(&resId, &resPassword)
	}
	if err != nil {
		_ = tx.Rollback()
		status.Error = "server error"
		status.Status = false
		statusSend, _ = proto.Marshal(&status)
		_ = req.GetConnection().SendMsg(0x103, statusSend)
		return
	}
	if resId == int(id) && resPassword == pwdMd5 {
		token, err := utils.CreateToken(&userLoginInfo)
		if err != nil {
			status.Error = err.Error()
			status.Status = false
			statusSend, _ = proto.Marshal(&status)
			_ = req.GetConnection().SendMsg(0x103, statusSend)
			_ = tx.Rollback()
			return
		}
		_ = tx.Commit()
		// 对token进行序列化后，发送给客户端
		tokenPro := protos.Token{Token: []byte(token)}
		tokenSend, _ := proto.Marshal(&tokenPro)
		color.Green("token byte len:%d", len(tokenSend))
		_ = req.GetConnection().SendMsg(0x301, tokenSend)
		// 将映射关系写入全局变量
		utils.ConnectionIdReflectorZinxConnID[int32(req.GetConnection().GetConnID())] = id
		// 将id写入属性
		req.GetConnection().SetProperty("id", id)
	} else {
		// id或密码不匹配
		status.Error = "id or password incorrect."
		status.Status = false
		statusSend, _ = proto.Marshal(&status)
		err = req.GetConnection().SendMsg(0x103, statusSend)
	}
}
