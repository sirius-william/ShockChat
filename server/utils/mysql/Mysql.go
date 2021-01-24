// @Title  Mysql.go
// @Description	数据库操作
// @Author  SiriusWilliam  2021-01-22 20:38
// @Update  2021-01-22 20:38
package mysql

type MySQLConfig struct {
	Host     string `json:"mysql.host"`
	Password string `json:"mysql.password"`
	Port     int    `json:"mysql.port"`
	Database string `json:"mysql.database"`
	Charset  string `json:"mysql.charset"`
}

var MysqlConfig MySQLConfig = MySQLConfig{
	Host: "127.0.0.1",
	Password: "",
	Port: 3306,
	Database: "shockchat",
	Charset: "utf8mb4",
}