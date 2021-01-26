// @Title  Register.go
// @Description	用户注册模块，消息id：
// @Author  SiriusWilliam  2021-01-22 9:50
// @Update  2021-01-22 9:50
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
	username := string(userNameBytes)
	password := string(passwordBytes)
	tel := string(telBytes)
	email := string(emailBytes)
	// 开始操作数据库
	// 开始事务
	tx, err := mysql.Mysql.Begin()
	if err != nil {
		color.Red(err.Error())
		return
	}
	// 拿到一个空闲的id
	/*
		 * SQL说明：该查询语句的执行策略如下
		+----+-------------+------------+------------+------+---------------+------------+---------+-------+------+----------+--------------------------+
		| id | select_type | table      | partitions | type | possible_keys | key        | key_len | ref   | rows | filtered | Extra                    |
		+----+-------------+------------+------------+------+---------------+------------+---------+-------+------+----------+--------------------------+
		|  1 | SIMPLE      | tb_user_id | NULL       | ref  | idx_status    | idx_status | 1       | const |   81 |   100.00 | Using where; Using index |
		+----+-------------+------------+------------+------+---------------+------------+---------+-------+------+----------+--------------------------+
		 * 以上可知，此查询语句走辅助索引idx_status，且不需要回表。锁：由for update加行级锁（因为走了索引，所以不是表锁），对查询到的行加写锁，禁止读。
		 * 行写锁有效保证了并发性能。
	*/
	row := tx.QueryRow("select id from tb_user_id where status = 1 limit 1 for update;")
	var id int
	err = row.Scan(&id)
	if err != nil {
		color.Red(err.Error())
		_ = tx.Rollback()
		return
	}
	// 修改id状态
	/*
	 * SQL说明：
	 */
	_, err = tx.Exec("update tb_user_id set status = 2 where id = ?;", id)
	if err != nil {
		color.Red(err.Error())
		_ = tx.Rollback()
		return
	}
	// 将用户信息写入表
	/*
	 * SQL说明：不走任何索引，为插入整行数据，下insert同
	 */
	_, err = tx.Exec("insert into tb_user_info(userid, username, tel, email) values (?, ?, ?, ?);", id, username, tel, email)
	if err != nil {
		color.Red(err.Error())
		_ = tx.Rollback()
		return
	}
	if err != nil {
		color.Red(err.Error())
		_ = tx.Rollback()
		return
	}
	// 将密码的md5写入表
	_, err = tx.Exec("insert into tb_user_login(userid, password) values(?, ?);", id, utils.StringToMD5(password))
	if err != nil {
		color.Red(err.Error())
		_ = tx.Rollback()
		return
	}
	err = tx.Commit()
	if err != nil {
		color.Red(err.Error())
		_ = tx.Rollback()
		return
	}
	userid := protos.UserId{Id: int64(id)}
	send, err := proto.Marshal(&userid)
	_ = req.GetConnection().SendMsg(0x201, send)
}
