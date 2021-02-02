// @Title  Friend.go
// @Description	获取好友列表的路由
// @Author  SiriusWilliam  2021-01-31 16:04
// @Update  2021-01-31 16:04
package router

import (
	"ShockChatServer/utils/mysql"
	"database/sql"
	"errors"
	"github.com/aceld/zinx/ziface"
	"github.com/aceld/zinx/znet"
)

/*
 * 消息Id：0x400
 * 回复：0x401
 */
type GetFriendListRouter struct {
	znet.BaseRouter
}

func (g *GetFriendListRouter) Handle(req ziface.IRequest) {

}

/*
 * 获取好友列表
 * 返回一个map，map字段为{id: "xxx", name: "xxx"}
 */
func GetFriendList(id string) (map[string]string, error) {
	// 返回结果集
	res := make(map[string]string)
	// 开启事务
	tx, err := mysql.Mysql.Begin()
	/*
		SQL：select userid, username from tb_user_info a inner join tb_friend b on b.id1=10000000 and b.id2=a.userid;
		该SQL执行计划为：
		+----+-------------+-------+------------+--------+---------------+---------+---------+-----------------+------+----------+-------------+
		| id | select_type | table | partitions | type   | possible_keys | key     | key_len | ref             | rows | filtered | Extra       |
		+----+-------------+-------+------------+--------+---------------+---------+---------+-----------------+------+----------+-------------+
		|  1 | SIMPLE      | b     | NULL       | ref    | idx_id1       | idx_id1 | 4       | const           |    3 |   100.00 | Using index |
		|  1 | SIMPLE      | a     | NULL       | eq_ref | PRIMARY       | PRIMARY | 4       | shockchat.b.id2 |    1 |   100.00 | NULL        |
		+----+-------------+-------+------------+--------+---------------+---------+---------+-----------------+------+----------+-------------+
		在执行内连接查询时，均走了索引，性能很高。
	*/

	var rows *sql.Rows
	rows, err = tx.Query("select userid, username from tb_user_info a inner join tb_friend b on b.id1=? and b.id2=a.userid;", id)
	err = tx.Commit()
	var userid string
	var username string
	for rows.Next() {
		err = rows.Scan(&userid, &username)
		if err != nil {
			break
		}
		res[userid] = username
	}
	if err != nil {
		return nil, err
	}
	return res, nil
}

/*
 * 添加好友
 * id1：发起添加好友请求的id
 * id2：被添加为好友的id
 * 说明：id1与id2不能为空，且不能相同
 */
func AddFriend(id1 int, id2 int) (int64, error) {
	if id1 == id2 || id1 <= 0 || id2 <= 0 {
		return -1, errors.New("id1 == id2 or one of them is empty")
	}
	var tx *sql.Tx
	var err error
	var res sql.Result
	tx, err = mysql.Mysql.Begin()
	res, err = tx.Exec("insert into tb_friend(id1, id2) values(?, ?);", id1, id2)
	if err != nil {
		_ = tx.Rollback()
		return -1, err
	}
	var affected int64
	affected, err = res.RowsAffected()
	if err == nil {
		return -1, err
	}
	return affected, err
}

/*
 * 删除好友
 * id1为发起删除好友的id
 * id2为被删除的id
 */
func DeleteFriend(id1 int, id2 int) (int64, error) {
	if id1 == id2 || id1 <= 0 || id2 <= 0 {
		return -1, errors.New("id1 == id2 or one of them is empty")
	}
	var tx *sql.Tx
	var err error
	var res sql.Result
	tx, err = mysql.Mysql.Begin()
	res, err = tx.Exec("delete from tb_friend where id1 = ? and id2 = ?;", id1, id2)
	if err != nil {
		_ = tx.Rollback()
		return -1, err
	}
	var affected int64
	affected, err = res.RowsAffected()
	if err == nil {
		return -1, err
	}
	return affected, err
}

type UserInfo struct {
	Id    int
	Name  string
	Tel   string
	Email string
}

/*
 * 查找好友
 */
func SearchFriend(id int) (*UserInfo, error) {
	if id <= 0 {
		return nil, errors.New("id must be greater than 0")
	}
	var res *UserInfo = &UserInfo{}
	var tx *sql.Tx
	var err error
	var rows *sql.Rows
	tx, err = mysql.Mysql.Begin()
	rows, err = tx.Query("select userid, username, tel, email from tb_user_info where userid=?", id)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		err = rows.Scan(&res.Id, &res.Name, &res.Tel, &res.Email)
		if err != nil {
			_ = tx.Rollback()
			return nil, err
		}
	}
	_ = tx.Commit()
	return res, nil
}
