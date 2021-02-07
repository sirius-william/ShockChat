// @Title  UserLoginDB.go
// @Description
// @Author  SiriusWilliam  2021-02-04 16:26
// @Update  2021-02-04 16:26
package mysql

import (
	"ShockChatServer/utils"
	"errors"
)

func UserLogin(userid int32, password string) (bool, error) {
	var err error
	tx, err := Mysql.Begin()
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
