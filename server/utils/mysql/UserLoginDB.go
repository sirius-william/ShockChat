// @Title  UserLoginDB.go
// @Description
// @Author  SiriusWilliam  2021-02-04 16:26
// @Update  2021-02-04 16:26
package mysql

import (
	"ShockChatServer/logger"
	"ShockChatServer/utils"
	"database/sql"
	"errors"
)

/*
 * 用户登录
 */
func UserLogin(userid int32, password string) (bool, error) {
	var err error
	tx, err := Mysql.Begin()
	// 检查用户账号状态，如果被冻结或不存在，则返回false
	isUserFrozen, err := tx.Query("select status from tb_user_id where id = ?;", userid)
	var status int32 = 0
	for isUserFrozen.Next() {
		err = isUserFrozen.Scan(&status)
		if err != nil {
			_ = tx.Rollback()
			return false, err
		}
	}
	switch status {
	case 0:
		{
			_ = tx.Commit()
			return false, errors.New("user not exist")
		}
	case 2:
		{
			_ = tx.Commit()
			return false, errors.New("user has been frozen")
		}
	}
	/*
		 * SQL说明：由MySQL评估，该查询语句策略为：
		+----+-------------+---------------+------------+-------+---------------+---------+---------+-------+------+----------+-------+
		| id | select_type | table         | partitions | type  | possible_keys | key     | key_len | ref   | rows | filtered | Extra |
		+----+-------------+---------------+------------+-------+---------------+---------+---------+-------+------+----------+-------+
		|  1 | SIMPLE      | tb_user_login | NULL       | const | PRIMARY       | PRIMARY | 4       | const |    1 |   100.00 | NULL  |
		+----+-------------+---------------+------------+-------+---------------+---------+---------+-------+------+----------+-------+
		由上，该查询语句走主键聚簇索引，性能卓越
	*/
	rows, err := tx.Query("select userid, password from tb_user_login where userid = ?;", userid)
	var resId int32
	var resPassword string
	for rows.Next() {
		err = rows.Scan(&resId, &resPassword)
		if err != nil {
			_ = tx.Rollback()
			return false, err
		}
	}
	if resId == userid && resPassword == utils.Password(password) {
		_ = tx.Commit()
		return true, nil
	}
	return false, errors.New("userid or password is incorrect")
}

/*
 * 操作指定账户的状态，status与mysql中的对应
 */
func FreezeUser(userid int32, status int32) (bool, error) {
	tx, err := Mysql.Begin()
	if err != nil {
		_ = tx.Rollback()
		return false, err
	}
	/*
		 * SQL策略
		 +----+-------------+------------+------------+-------+---------------+---------+---------+-------+------+----------+-------------+
		| id | select_type | table      | partitions | type  | possible_keys | key     | key_len | ref   | rows | filtered | Extra       |
		+----+-------------+------------+------------+-------+---------------+---------+---------+-------+------+----------+-------------+
		|  1 | UPDATE      | tb_user_id | NULL       | range | PRIMARY       | PRIMARY | 4       | const |    1 |   100.00 | Using where |
		+----+-------------+------------+------------+-------+---------------+---------+---------+-------+------+----------+-------------+
		走主键聚簇索引，性能极高
	*/
	res, err := tx.Exec("update tb_user_id set status = ? where id = ?;", status, userid)
	if err != nil {
		_ = tx.Rollback()
		return false, err
	}
	if affect, err := res.RowsAffected(); affect == 1 {
		err = tx.Commit()
		if err != nil {
			logger.Log.Error(err.Error())
			return false, err
		}
		return true, nil
	} else {
		_ = tx.Rollback()
		return false, err
	}
}

const (
	NormalNotSignedIn = iota // 正常，未登录
	NormalSignedIn           // 正常，已登录
	Frozen                   // 冻结
	NonExisted               // 账户不存在
)

// 查看id的状态
func GetUserStatus(userid int32) (int32, error) {
	var err error
	tx, err := Mysql.Begin()
	if err != nil {
		logger.Log.Error(err.Error())
		return -1, err
	}
	var rows *sql.Rows
	rows, err = tx.Query("select status from tb_user_id where id = ? for update;", userid)
	if err != nil {
		_ = tx.Rollback()
		return -1, err
	}
	var status int32
	var count int = 0
	var res int32
	for rows.Next() {
		err = rows.Scan(&status)
		if err != nil {
			_ = rows.Close()
			_ = tx.Rollback()
			return -1, err
		}
		count += 1
	}
	if count == 1 {
		switch status {
		case 1:
			res = NonExisted
		case 2:
			{
				if utils.ConnectionIdReflectorZinxConnID.IsLogin(userid) {
					res = NormalSignedIn
				} else {
					res = NormalNotSignedIn
				}
			}
		case 3:
			res = Frozen
		}
	} else if count == 0 {
		res = NonExisted
	} else {
		return -1, errors.New("The Count does not match one. ")
	}
	return res, nil
}

/*
 * 清除指定账户的所有资产，包括消息记录
 */
func ClearUserProperty(userid int32) {

}
