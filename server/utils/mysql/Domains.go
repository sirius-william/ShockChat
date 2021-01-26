// @Title  Domains.go
// @Description	一些实体类（结构体），对应MySQL表
// @Author  SiriusWilliam  2021-01-26 17:19
// @Update  2021-01-26 17:19
package mysql

/*
 * for tb_user_info
 */
type UserInfo struct {
	UserId   int    `db:"userid"`
	Username string `db:"username"`
	Tel      string `db:"tel"`
	Email    string `db:"email"`
}

/*
 * for tb_user_login
 */
type UserLogin struct {
	Id       int    `db:"userid"`
	Password string `db:"password"`
}

/*
 * for tb_friend
 */
type Friend struct {
	Id  int `db:"id"`
	Id1 int `db:"id1"`
	Id2 int `db:"id2"`
}

/*
 * for tb_user_id
 */
type UserId struct {
	Id     int `db:"userid"`
	Status int `db:"status"`
}
