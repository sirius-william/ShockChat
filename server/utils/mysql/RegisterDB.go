// @Title  RegisterDB.go
// @Description
// @Author  SiriusWilliam  2021-02-04 20:43
// @Update  2021-02-04 20:43
package mysql

import (
	"ShockChatServer/utils"
)

func Register(username string, password string, tel string, email string) (int, error) {
	var err error
	tx, err := Mysql.Begin()
	if err != nil {
		return -1, err
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
	row := tx.QueryRow("select id from tb_user_id where status = 1 limit 1 for update skip locked;")
	var id int
	err = row.Scan(&id)
	if err != nil {
		_ = tx.Rollback()
		return -1, err
	}
	// 修改id状态
	_, err = tx.Exec("update tb_user_id set status = 2 where id = ?;", id)
	if err != nil {
		_ = tx.Rollback()
		return -1, err
	}
	// 将用户信息写入表
	/*
	 * SQL说明：不走任何索引，为插入整行数据，下insert同
	 */
	_, err = tx.Exec("insert into tb_user_info(userid, username, tel, email) values (?, ?, ?, ?);", id, username, tel, email)
	if err != nil {
		_ = tx.Rollback()
		return -1, err
	}
	if err != nil {
		_ = tx.Rollback()
		return -1, err
	}
	// 将密码的md5写入表
	_, err = tx.Exec("insert into tb_user_login(userid, password) values(?, ?);", id, utils.Password(password))
	if err != nil {
		_ = tx.Rollback()
		return -1, err
	}
	err = tx.Commit()
	if err != nil {
		_ = tx.Rollback()
		return -1, err
	}
	return id, nil
}
